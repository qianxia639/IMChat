package service

import (
	db "IMChat/db/sqlc"
	errDefine "IMChat/internal/errors"
	"IMChat/internal/validator"
	"IMChat/pb"
	"IMChat/utils"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	*Server
}

func NewUserService(server *Server) pb.UserServiceServer {
	return &UserService{
		Server: server,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	createUserValidator := &validator.CreateUserValidator{}
	if err := validator.NewValidateContext(createUserValidator).Validate(req); err != nil {
		log.Err(err).Msg("Create User parmas validator")
		return nil, err
	}

	// 校验邮箱验证码
	// ok, _ := email.VerifyEmailCode(userService.cache, req.Email, req.EmailCode)
	// if !ok {
	// 	return nil, errDefine.EmailCodeErr
	// }

	hashPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		return nil, errDefine.ParamsErr
	}

	// 判断用户是否存在
	if user, _ := userService.store.GetUser(ctx, req.Username); user.ID > 0 {
		return nil, errDefine.AccountExistErr
	}

	// 判断邮箱是否存在
	if i, _ := userService.store.ExistEmail(ctx, req.Email); i > 0 {
		return nil, errDefine.EmailExistErr
	}

	gender := req.GetGender()
	if gender == pb.Gender_UNKNOWN {
		gender = 3
	}

	arg := &db.CreateUserParams{
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
		Nickname: req.Username,
		Gender:   int16(gender),
	}

	_, err = userService.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, errDefine.ServerErr
	}

	return &pb.CreateUserResponse{
		Message: "Create User Succeddfully",
	}, nil
}

func (userService *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	// 参数校验
	loginUserValidator := &validator.LoginUserValidator{}
	if err := validator.NewValidateContext(loginUserValidator).Validate(req); err != nil {
		return nil, errDefine.ParamsErr
	}

	// 校验账户是否锁定
	lockedUsernameKey := getAccountLocked(time.Now().Format("2006-01-02T15:00:00"), req.Username)
	isLocked, err := userService.cache.Get(ctx, lockedUsernameKey).Bool()
	if err != nil && err != redis.Nil {
		return nil, errDefine.ServerErr
	}

	if isLocked {
		return nil, errDefine.AccountLockedErr
	}

	// 获取用户信息
	user, err := userService.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			log.Error().Err(err).Msgf("GetUser(_)")
			return nil, errDefine.UserNotFoundErr
		}
		return nil, errDefine.ServerErr
	}

	loginAttemptsKey := getLoginAttempts(req.Username)
	// 校验密码
	err = utils.Decrypt(req.Password, user.Password)
	if err != nil {
		if err := userService.recordLoginAttempts(ctx, loginAttemptsKey, lockedUsernameKey); err != nil {
			return nil, err
		}
		return nil, errDefine.IncorrectPasswordErr
	}

	// 重置失败次数
	err = userService.cache.Del(ctx, loginAttemptsKey).Err()
	if err != nil {
		return nil, errDefine.ServerErr
	}

	// 查询最新登录日志
	userLoginLog, _ := userService.store.GetLastUserLoginLog(ctx, user.ID)

	isLoginExcptional := false
	mtdt := extractMetadata(ctx)

	// 异地登录
	// TODO: 出现异地登录时需要发送消息告知用户(系统消息发送)
	if userLoginLog.ID > 0 && userLoginLog.LoginIp != mtdt.ClientIp {
		log.Print("异地登录,非本人操作请及时修改密码")
		isLoginExcptional = true
	}

	// 创建token
	accessToken, err := userService.maker.CreateToken(user.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	region, _ := utils.GetRegion(mtdt.ClientIp)

	// 添加登录日志
	_, err = userService.store.AddUserLoginLog(ctx, &db.AddUserLoginLogParams{
		UserID:             user.ID,
		LoginIp:            mtdt.ClientIp,
		LoginIpRegion:      region,
		IsLoginExceptional: isLoginExcptional,
		Platform:           "",
		UserAgent:          mtdt.UserAgent,
	})
	if err != nil {
		log.Warn().Msgf("failed to add user login log: %v", err)
	}

	// 更改在线状态
	now := time.Now()
	arg := &db.UpdateUserParams{
		Username:  user.Username,
		UpdatedAt: now,
		OnlineStatus: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	}
	user, err = userService.store.UpdateUser(ctx, arg)
	if err != nil {
		return nil, errDefine.ServerErr
	}

	expireAt := time.Duration(utils.RandomInt(27, 30))
	if err := userService.cache.Set(ctx, getUserInfoKey(user.Username), &user, expireAt*time.Minute).Err(); err != nil {
		log.Error().Err(err).Msgf("LoginUser HSET error")
		return nil, errDefine.ServerErr
	}

	return &pb.LoginUserResponse{
		AccessToken: accessToken,
	}, nil
}

// 记录登录失败次数
func (userService *UserService) recordLoginAttempts(ctx context.Context, loginAttemptsKey, lockedUsernameKey string) error {

	err := userService.cache.SetNX(ctx, loginAttemptsKey, 0, time.Hour).Err()
	log.Error().Err(err).Msgf("record login attempts setnx error")

	// 累加失败次数
	if err := userService.cache.Incr(ctx, loginAttemptsKey).Err(); err != nil {
		return errDefine.ServerErr
	}
	// 获取失败登录次数
	loginAttempts, err := userService.cache.Get(ctx, loginAttemptsKey).Int()
	if err != nil && err != redis.Nil {
		return errDefine.ServerErr
	}

	// 判断是否大于最大失败次数
	if loginAttempts >= 5 {
		// 锁定账户
		err := userService.cache.Set(ctx, lockedUsernameKey, true, time.Hour).Err()
		if err != nil {
			return errDefine.ServerErr
		}
		return errDefine.AccountLockedErr
	}
	return nil
}

func (userService *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := userService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// var user db.User
	// _ = userService.cache.Get(ctx, getUserInfoKey(req.Username)).Scan(&user)

	if user.ID <= 0 {
		*user, _ = userService.store.GetUser(ctx, req.Username)
	}

	resp := &pb.GetUserResponse{
		User: converUser(*user),
	}

	return resp, nil
}

// TODO 该接口后期需要进行重构
func (userService *UserService) SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.SearchUserResponse, error) {
	switch req.GetType() {
	case 1:
		user, _ := userService.store.GetUser(ctx, req.GetQuery())
		usereAny, err := anypb.New(converUser(user))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "user marshal: %v", err)
		}
		return &pb.SearchUserResponse{Data: usereAny}, nil
	case 2:
		log.Info().Msg("Search cluster...")
		return nil, nil
	default:
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// return &pb.SearchUserResponse{}, nil
}

func (userService *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	user, err := userService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.Username != req.Username {
		log.Error().Msg("cannot update other usre's info")
		return nil, errDefine.PermissionDeniedErr
	}

	if req.Nickname != nil {
		if i, _ := userService.store.ExistNickname(ctx, *req.Nickname); i > 0 {
			return nil, errDefine.NicknameExistErr
		}
	}

	if req.Gender != nil {
		_, exists := pb.Gender_name[int32(*req.Gender)]
		if !exists {
			// return nil, status.Errorf(codes.InvalidArgument, "invlaid argument, gender is %d", req.GetGender())
			log.Error().Msgf("invalid argument, gender is %d", int32(*req.Gender))
			return nil, errDefine.ParamsErr
		}
	}

	arg := &db.UpdateUserParams{
		Username:  req.GetUsername(),
		UpdatedAt: time.Now(),
		Nickname: pgtype.Text{
			String: req.GetNickname(),
			Valid:  req.Nickname != nil,
		},
		Gender: pgtype.Int2{
			Int16: int16(req.GetGender()),
			Valid: req.Gender != nil,
		},
	}

	*user, err = userService.store.UpdateUser(ctx, arg)
	if err != nil {
		log.Error().Err(err).Msg("Can't update user")
		return nil, errDefine.ServerErr
	}

	userInfoKey := getUserInfoKey(user.Username)
	expireAt := time.Duration(utils.RandomInt(27, 30))
	if err := userService.cache.Set(ctx, userInfoKey, &user, expireAt*time.Minute).Err(); err != nil {
		return nil, errDefine.ServerErr
	}

	resp := &pb.UpdateUserResponse{
		User: converUser(*user),
	}

	return resp, nil
}

func (userService *UserService) DeleteUser(ctx context.Context, req *emptypb.Empty) (*pb.DeleteUserResponse, error) {
	user, err := userService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// userInfoKey := getUserInfoKey(payload.Username)
	// var user db.User
	// if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
	// 	return nil, errDefine.ServerErr
	// }

	if err := userService.store.DeleteUserTx(ctx, user.ID); err != nil {
		return nil, errDefine.ServerErr
	}

	userInfoKey := getUserInfoKey(user.Username)
	_ = userService.cache.Del(ctx, userInfoKey).Err()

	resp := &pb.DeleteUserResponse{
		Message: "Delete Successfully...",
	}

	return resp, nil
}

func (userService *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {

	user, err := userService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// TODO 校验邮箱验证码是否正确
	if req.EmailCode != "123" {
		return nil, errDefine.EmailCodeErr
	}

	// 校验两次密码是否一致
	if req.UserPassword != req.UserConfirmPassword {
		return nil, errDefine.RestPwdNotSame
	}

	// 判断新旧密码是否一致
	if err := utils.Decrypt(req.UserConfirmPassword, user.Password); err == nil {
		return nil, errDefine.NewPwdIsSame
	}

	// 校验上次密码更新时间与当前时间差
	if !time.Now().After(user.PasswordChangedAt.Add(7 * 24 * time.Hour)) {
		return nil, errDefine.NewPwdBelowInterval
	}

	hashPassword, err := utils.Encrypt(req.GetUserConfirmPassword())
	if err != nil {
		return nil, errDefine.ParamsErr
	}

	now := time.Now()
	arg := &db.UpdateUserPasswordParams{
		Password:          hashPassword,
		PasswordChangedAt: now,
		Username:          user.Username,
		UpdatedAt:         now,
	}

	err = userService.store.UpdateUserPassword(ctx, arg)
	if err != nil {
		return nil, errDefine.ServerErr
	}

	_ = userService.cache.Del(ctx, getUserInfoKey(user.Username)).Err()

	resp := &pb.UpdateUserPasswordResponse{
		Message: "Update Successfully...",
	}

	return resp, nil
}

func (userService *UserService) Logout(ctx context.Context, req *emptypb.Empty) (*pb.LogoutResponse, error) {
	user, err := userService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// 更改在线状态为离线并更新最后在线时间
	now := time.Now()
	arg := &db.UpdateUserParams{
		Username:  user.Username,
		UpdatedAt: now,
		OnlineStatus: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
		LastLoginAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}
	_, err = userService.store.UpdateUser(ctx, arg)
	if err != nil {
		return nil, errDefine.ServerErr
	}

	userInfoKey := getUserInfoKey(user.Username)
	_ = userService.cache.Del(ctx, userInfoKey).Err()

	resp := &pb.LogoutResponse{
		Message: "Logout Successfully...",
	}

	return resp, nil
}
