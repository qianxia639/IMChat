package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	to := "example@example.com"
	subject := "使用Golang发送邮件"
	body := "测试邮件"

	err := SendMail(to, subject, body)
	require.NoError(t, err)
}
