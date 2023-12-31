package validator

import (
	"IMChat/internal/errors"
	"IMChat/pb"
	"IMChat/utils"
	"fmt"
	"regexp"

	"github.com/dlclark/regexp2"
	"github.com/rs/zerolog/log"
)

var (
	isValidUsername = regexp.MustCompile(`^[\da-zA-Z_]{6,20}$`).MatchString
	passwordRegexp  = "^(?=.*[A-Za-z])(?=.*\\d|.*[\\W_])[A-Za-z\\d\\W_]{8,16}$"
)

type LoginUserValidator struct{}

var _ Validator = (*LoginUserValidator)(nil)

func (*LoginUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.LoginUserRequest)
	if !ok {
		return fmt.Errorf("interface{} type cannot be converted to *pb.LoginUserRequest type")
	}

	// 校验用户名
	if !isValidUsername(req.Username) {
		return errors.UsernameFormatErr
	}
	// 校验密码
	reg, _ := regexp2.Compile(passwordRegexp, 0)
	_, err := reg.MatchString(req.Password)
	if err != nil {
		return errors.PasswordFormatErr
	}

	return nil
}

type CreateUserValidator struct{}

var _ Validator = (*CreateUserValidator)(nil)

func (*CreateUserValidator) Validate(param interface{}) error {
	req, ok := param.(*pb.CreateUserRequest)
	if !ok {
		return fmt.Errorf("interface{} type cannot be converted to *pb.CreateUserRequest type")
	}

	// 校验用户名
	if !isValidUsername(req.Username) {
		return errors.UsernameFormatErr
	}

	// 校验密码
	reg, _ := regexp2.Compile(passwordRegexp, 0)
	_, err := reg.MatchString(req.Password)
	if err != nil {
		log.Error().Err(err).Msgf("密码格式匹配错误")
		return errors.PasswordFormatErr
	}

	err = utils.ValidatorEmail(req.Email)
	if err != nil {
		log.Error().Err(err).Msgf("is not a valid email address")
		return errors.InvalidEmailAddrErr
	}

	// 判断gener
	if _, ok := pb.Gender_name[int32(req.Gender)]; !ok {
		return fmt.Errorf("invalid gender")
	}

	return nil
}
