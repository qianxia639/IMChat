package validator

import (
	"IMChat/pb"
	"IMChat/utils"
	"fmt"
	"regexp"

	"github.com/dlclark/regexp2"
)

var (
	isValidUsername = regexp.MustCompile(`^[\da-zA-Z_]{6,20}$`).MatchString
	isValidPassword = regexp.MustCompile(`^[a-zA-Z\d]{8,16}$`).MatchString
)

type LoginUserValidator struct{}

var _ Validator = (*LoginUserValidator)(nil)

func (*LoginUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.LoginUserRequest)
	if !ok {
		return fmt.Errorf("invalid argument")
	}

	// 校验用户名
	if !isValidUsername(req.Username) {
		return fmt.Errorf("invalid username")
	}
	// 校验密码
	reg, _ := regexp2.Compile("^(?=.*[A-Za-z])(?=.*\\d|.*[\\W_])[A-Za-z\\d\\W_]{8,16}$", 0)
	_, err := reg.MatchString(req.Password)
	if err != nil {
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
	if !isValidUsername(req.Username) {
		return fmt.Errorf("invalid username")
	}
	// 校验密码
	if !isValidPassword(req.Password) {
		return fmt.Errorf("invalid password")
	}

	err := utils.ValidatorEmail(req.Email)
	if err != nil {
		return err
	}

	return nil
}
