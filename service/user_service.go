package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"IMChat/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"

	"github.com/lib/pq"
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

	// 获取用户信息
	user, err := userService.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	// 校验密码
	err = utils.Decrypt(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "incorrect password:")
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
		// log.Printf("failed to add user login log: %v", err)
		log.Warn().Msgf("failed to add user login log: %v", err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", user.Username)
	expireAt := time.Duration(utils.RandomInt(27, 30))
	if err := userService.cache.Set(ctx, userInfoKey, &user, expireAt*time.Minute).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set user info: %v", err)
	}

	return &pb.LoginUserResponse{
		AccessToken: accessToken,
	}, nil
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %v", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &pb.CreateUserResponse{
		Message: "Create User Succeddfully",
	}, nil
}

func (userService *UserService) GetUser(ctx context.Context, req *pb.EmptyRequest) (*pb.GetUserResponse, error) {

	payload, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)

	var user db.User
	if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	resp := &pb.GetUserResponse{
		User: converUser(user),
	}

	return resp, nil
}

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
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", user.Username)
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

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
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
	payload, err := userService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// TODO 校验邮箱验证码是否正确
	if req.EmailCode != "123" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	}

	// 校验两次密码是否一致
	if req.UserPassword != req.UserConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "两次密码不一致")
	}

	// 判断新旧密码是否一致
	if err := utils.Decrypt(req.UserConfirmPassword, user.Password); err == nil {
		return nil, status.Error(codes.InvalidArgument, "新旧密码不能一致")
	}

	// userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	// var user db.User
	// if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	// }

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

	_ = userService.cache.Del(ctx, userInfoKey).Err()

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

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	if err := userService.cache.Get(ctx, userInfoKey).Scan(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	}

	// _ = userService.store.UpdateLastUserLoginLog(ctx, sql.NullInt32{Int32: user.ID, Valid: true})

	_ = userService.cache.Del(ctx, userInfoKey).Err()

	resp := &pb.LogoutResponse{
		Message: "Logout Successfully...",
	}

	return resp, nil
}
