package token

import (
	"IMChat/core/config"
	"IMChat/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	tokenConf := config.Token{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	maker := NewPasetoMaker(tokenConf)

	username := utils.RandomString(6)

	issuedAt := time.Now()
	expiredAt := time.Now().Add(tokenConf.AccessTokenDuration)

	token, err := maker.CreateToken(username)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)

	t.Logf("payload: %+v\n", payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpireAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	tokenConf := config.Token{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: -time.Minute,
	}
	maker := NewPasetoMaker(tokenConf)

	token, err := maker.CreateToken(utils.RandomString(6))
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
