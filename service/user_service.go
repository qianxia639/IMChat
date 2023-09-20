package service

import (
	db "IMChat/db/sqlc"
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
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Server
}

func NewUserService(server Server) pb.UserServiceServer {
	return &UserService{
		Server: server,
	}
}

func (userService *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	// 参数校验
	loginUserValidator := &validator.LoginUserValidator{}
	if err := loginUserValidator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// 校验账户是否锁定
	lockedUsernameKey := getLocked(time.Now().Format("2006-01-02 15:00:00"), req.GetUsername())
	isLocked, err := userService.cache.Get(ctx, lockedUsernameKey).Bool()
	if err != nil && err != redis.Nil {
		return nil, status.Errorf(codes.Internal, "failed to check user lock status: %v\n", err)
	}

	if isLocked {
		return nil, status.Errorf(codes.PermissionDenied, "account locked, please try again later")
	}

	// 获取用户信息
	user, err := userService.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	loginAttemptsKey := getLoginAttempts(req.GetUsername())
	// 校验密码
	err = utils.Decrypt(req.GetPassword(), user.Password)
	if err != nil {
		if err := userService.recordLoginAttempts(ctx, loginAttemptsKey, lockedUsernameKey); err != nil {
			return nil, err
		}
		return nil, status.Error(codes.Unauthenticated, "incorrect password")
	}

	// 重置失败次数
	err = userService.cache.Del(ctx, loginAttemptsKey).Err()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to reset login attempts: %v", err)
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

	expireAt := time.Duration(utils.RandomInt(27, 30))
	if err := userService.cache.Set(ctx, getUserInfoKey(user.Username), &user, expireAt*time.Minute).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set user info: %v", err)
	}

	return &pb.LoginUserResponse{
		AccessToken: accessToken,
	}, nil
}

// 记录登录失败次数
func (userService *UserService) recordLoginAttempts(ctx context.Context, loginAttemptsKey, lockedUsernameKey string) error {
	// 累加失败次数
	if err := userService.cache.Incr(ctx, loginAttemptsKey).Err(); err != nil {
		return status.Errorf(codes.Internal, "failed to increment login attempts: %v", err)
	}
	// 获取失败登录次数
	loginAttempts, err := userService.cache.Get(ctx, loginAttemptsKey).Int()
	if err != nil && err != redis.Nil {
		return status.Errorf(codes.Internal, "failed to get login attempts: %v", err)
	}

	// 判断是否大于最大失败次数
	if loginAttempts >= 5 {
		// 锁定账户
		err := userService.cache.Set(ctx, lockedUsernameKey, true, time.Hour).Err()
		if err != nil {
			return status.Errorf(codes.Internal, "failed to lock account: %v\n", err)
		}
		return status.Errorf(codes.PermissionDenied, "account locked, please try again later")
	}
	return nil
}

func (userService *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	err := utils.ValidatorEmail(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hashPassword, err := utils.Encrypt(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	if req.Gender != nil {
		_, exists := pb.Gender_name[int32(req.GetGender())]
		if !exists {
			return nil, status.Errorf(codes.InvalidArgument, "invlaid argument, gender is %d", req.GetGender())
		}
	}

	arg := &db.CreateUserParams{
		Username: req.GetUsername(),
		Password: hashPassword,
		Email:    req.GetEmail(),
		Nickname: req.GetUsername(),
		Gender:   int16(req.GetGender()),
	}

	_, err = userService.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &pb.CreateUserResponse{
		Message: "Create User Succeddfully",
	}, nil
}

// TODO 该接口后期需要优化
func (userService *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	_, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	var user db.User
	_ = userService.cache.Get(ctx, getUserInfoKey(req.GetUsername())).Scan(&user)

	if user.ID <= 0 {
		user, _ = userService.store.GetUser(ctx, req.Username)
	}

	resp := &pb.GetUserResponse{
		User: converUser(user),
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

	payload, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if payload.Username != req.GetUsername() {
		return nil, status.Error(codes.PermissionDenied, "cannot update other usre's info")
	}

	if req.Gender != nil {
		_, exists := pb.Gender_name[int32(req.GetGender())]
		if !exists {
			return nil, status.Errorf(codes.InvalidArgument, "invlaid argument, gender is %d", req.GetGender())
		}
	}

	arg := &db.UpdateUserParams{
		Username:  req.GetUsername(),
		UpdatedAt: time.Now(),
		Nickname: pgtype.Text{
			String: req.GetNickname(),
			Valid:  req.Nickname != nil,
		},
		Avatar: pgtype.Text{
			String: req.GetAvatar(),
			Valid:  req.Avatar != nil,
		},
		Gender: pgtype.Int2{
			Int16: int16(req.GetGender()),
			Valid: req.Gender != nil,
		},
	}

	user, err := userService.store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	userInfoKey := getUserInfoKey(payload.Username)
	expireAt := time.Duration(utils.RandomInt(27, 30))
	if err := userService.cache.Set(ctx, userInfoKey, &user, expireAt*time.Minute).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set user info: %v", err)
	}

	resp := &pb.UpdateUserResponse{
		User: converUser(user),
	}

	return resp, nil
}

func (userService *UserService) DeleteUser(ctx context.Context, req *pb.EmptyRequest) (*pb.DeleteUserResponse, error) {
	payload, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userInfoKey := getUserInfoKey(payload.Username)
	var user db.User
	if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	}

	if err := userService.store.DeleteUser(ctx, user.ID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	_ = userService.cache.Del(ctx, userInfoKey).Err()

	resp := &pb.DeleteUserResponse{
		Message: "Delete Successfully...",
	}

	return resp, nil
}

func (userService *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {

	user, err := userService.getUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	// TODO 校验邮箱验证码是否正确
	if req.EmailCode != "123" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	// 校验两次密码是否一致
	if req.UserPassword != req.UserConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "两次密码不一致")
	}

	// 判断新旧密码是否一致
	if err := utils.Decrypt(req.UserConfirmPassword, user.Password); err == nil {
		return nil, status.Error(codes.InvalidArgument, "新旧密码不能一致")
	}

	// 校验上次密码更新时间与当前时间差
	if !time.Now().After(user.PasswordChangedAt.Add(7 * 24 * time.Hour)) {
		return nil, status.Error(codes.InvalidArgument, "两次修改密码的间隔不得低于7天")
	}

	hashPassword, err := utils.Encrypt(req.GetUserConfirmPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
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
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	_ = userService.cache.Del(ctx, getUserInfoKey(user.Username)).Err()

	resp := &pb.UpdateUserPasswordResponse{
		Message: "Update Successfully...",
	}

	return resp, nil
}

func (userService *UserService) Logout(ctx context.Context, req *pb.EmptyRequest) (*pb.LogoutResponse, error) {
	payload, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userInfoKey := getUserInfoKey(payload.Username)
	var user db.User
	if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	}

	_ = userService.cache.Del(ctx, userInfoKey).Err()

	resp := &pb.LogoutResponse{
		Message: "Logout Successfully...",
	}

	return resp, nil
}
