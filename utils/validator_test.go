package utils

import "testing"

func TestValidatorEmail(t *testing.T) {
	emails := []string{
		"dfjh@gmail.com",
		"dfkjl@163.com",
		"kfkh@outlook.com",
		"skdklkfj@qq.com",
		"@",
		"kdfh",
	}

	for _, v := range emails {
		_ = ValidatorEmail(v)

	}
}
