package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
	"github.com/Adetunjii/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"

)

func createRandomTransaction(t *testing.T) Transaction {
	arg := CreateTransactionDto{
		AccountID: 		 8,
		Amount:          util.RandomBalance(),
		TransactionType: "DEBIT",
		Currency:        util.RandomCurrency(),
		Status: 		"PENDING",
		Reference:       util.RandomString(10),
	}

	transaction, err := testRepository.CreateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	return transaction
}

func TestRepository_CreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestRepository_GetTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)

	transaction2, err := testRepository.GetTransaction(context.Background(), transaction1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transaction2)
	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.TransactionType, transaction2.TransactionType)
	require.Equal(t, transaction1.Currency, transaction2.Currency)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Status, transaction2.Status)
}


func TestRepository_ListTransactions(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransaction(t)
	}

	arg := ListTransactionParams{
		Limit:  5,
		Offset: 5,
	}

	transactions, err := testRepository.ListTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}


}

func TestRepository_UpdateTransaction(t *testing.T) {

	transaction1 := createRandomTransaction(t)

	args := UpdateTransactionDto{
		ID:     transaction1.ID,
		Status: "SUCCESSFUL",
	}

	transaction2, err := testRepository.UpdateTransaction(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, transaction1)
	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, args.Status, transaction2.Status)
	require.Equal(t, transaction1.Currency, transaction2.Currency)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
}

func TestRepository_DeleteTransaction(t *testing.T) {
	transaction := createRandomTransaction(t)

	err :=  testRepository.DeleteTransaction(context.Background(), transaction.ID)

	require.NoError(t, err)
}