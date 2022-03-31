package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const minSecretKeySize = 32

type JwtFactory struct {
	secretKey string
}

func NewJwtFactory(secretKey string) (TokenFactory, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JwtFactory{secretKey: secretKey}, nil
}

func (j *JwtFactory) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JwtFactory) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, InvalidTokenError
		}
		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		validationError, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationError.Inner, ExpiredJwtError) {
			return nil, ExpiredJwtError
		}
		return nil, InvalidTokenError
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, InvalidTokenError
	}
	return payload, nil
}
