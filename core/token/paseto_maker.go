package token

import (
	"IMChat/core/config"
	"IMChat/internal/bytesconv"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     time.Duration
	tokenConf    config.Token
}

func NewPasetoMaker(tokenConf config.Token) Maker {
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		// symmetricKey: []byte(symmetricKey),
		// duration:     duration,
		tokenConf: tokenConf,
	}
	return maker
}

func (maker *PasetoMaker) CreateToken(username string) (string, error) {
	if len(maker.tokenConf.TokenSymmetricKey) != chacha20poly1305.KeySize {
		return "", ErrKetTooShort
	}

	payload := NewPayload(username, maker.tokenConf.AccessTokenDuration)
	return maker.paseto.Encrypt([]byte(maker.tokenConf.TokenSymmetricKey), payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, bytesconv.StringToBytes(maker.tokenConf.TokenSymmetricKey), payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
