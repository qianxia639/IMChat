package validator

import (
	"IMChat/pb"
	"errors"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[\da-zA-Z_]{6,20}$`).MatchString
	isValidPassword = regexp.MustCompile(`^[\da-zA-Z_?!]{6,20}$`).MatchString
)

type LoginUserValidator struct{}

func (*LoginUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.LoginUserRequest)
	if !ok {
		return errors.New("invalid argument")
	}

	// 校验用户名
	if !isValidUsername(req.GetUsername()) {
		return errors.New("invalid username")
	}
	// 校验密码
	if !isValidPassword(req.GetPassword()) {
		return errors.New("invalid password")
	}
	return nil
}
