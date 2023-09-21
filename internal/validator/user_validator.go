package validator

import (
	"IMChat/pb"
	"IMChat/utils"
	"fmt"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[\da-zA-Z_]{6,20}$`).MatchString
	isValidPassword = regexp.MustCompile(`^[\da-zA-Z_?!]{6,20}$`).MatchString
)

type LoginUserValidator struct{}

var _ Validator = (*LoginUserValidator)(nil)

func (*LoginUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.LoginUserRequest)
	if !ok {
		return fmt.Errorf("invalid argument")
	}

	// 校验用户名
	if !isValidUsername(req.GetUsername()) {
		return fmt.Errorf("invalid username")
	}
	// 校验密码
	if !isValidPassword(req.GetPassword()) {
		return fmt.Errorf("invalid password")
	}
	return nil
}

type CreateUserValidator struct{}

var _ Validator = (*CreateUserValidator)(nil)

func (*CreateUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.CreateUserRequest)
	if !ok {
		return fmt.Errorf("invalid argument")
	}

	// 校验用户名
	if !isValidUsername(req.GetUsername()) {
		return fmt.Errorf("invalid username")
	}
	// 校验密码
	if !isValidPassword(req.GetPassword()) {
		return fmt.Errorf("invalid password")
	}

	err := utils.ValidatorEmail(req.GetEmail())
	if err != nil {
		return err
	}

	if req.Gender != nil {
		_, exists := pb.Gender_name[int32(req.GetGender())]
		if !exists {
			return fmt.Errorf("invlaid argument, gender is %d", req.GetGender())
		}
	}

	return nil
}
