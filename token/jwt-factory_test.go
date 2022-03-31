package token

import (
	"github.com/Adetunjii/simplebank/util"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJwtFactory(t *testing.T) {
	factory, err := NewJwtFactory(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwnerName()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := factory.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := factory.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	factory, err := NewJwtFactory(util.RandomString(32))
	require.NoError(t, err)

	token, err := factory.CreateToken(util.RandomOwnerName(), -time.Minute)
	require.NoError(t, err)

	payload, err := factory.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ExpiredJwtError.Error())
	require.Nil(t, payload)
}

func TestInvalidJwtAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomOwnerName(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	factory, err := NewJwtFactory(util.RandomString(32))
	require.NoError(t, err)

	decodedPayload, err := factory.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, InvalidTokenError.Error())
	require.Nil(t, decodedPayload)
}
