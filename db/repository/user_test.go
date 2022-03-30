package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
	"github.com/Adetunjii/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserDto{
		Username: util.RandomOwnerName(),
		Password: util.RandomString(10),
		FullName: util.RandomOwnerName(),
		Email:	util.RandomEmail(),
	}

	user, err := testRepository.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestRepository_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestRepository_GetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testRepository.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Email, user1.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
