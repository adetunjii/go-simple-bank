package db

import (
	"context"
	"database/sql"
	. "github.com/Adetunjii/simplebank/db/models"
	"github.com/Adetunjii/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountDto{
		Owner: util.RandomOwnerName(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testRepository.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestRepository_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestRepository_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testRepository.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestRepository_UpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	args := UpdateAccountDto{
		ID:      account1.ID,
		Balance: util.RandomBalance(),
	}

	account2, err := testRepository.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestRepository_DeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testRepository.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testRepository.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestRepository_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	
	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testRepository.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {

		require.NotEmpty(t, account)
	}

}