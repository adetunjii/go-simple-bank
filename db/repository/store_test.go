package db

import (
	"context"
	"github.com/Adetunjii/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := CreateNewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxnResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxnParams{
				SourceAccountID:      account1.ID,
				DestinationAccountID: account2.ID,
				Amount:               amount,
				Currency:             util.RandomCurrency(),
				Reference: 			util.RandomString(10),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++{
		err := <- errs
		require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.SourceAccountID)
		require.Equal(t, account2.ID, transfer.DestinationAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		sourceAccountTransaction := result.SourceTransaction
		require.NotEmpty(t, sourceAccountTransaction)
		require.Equal(t, "DEBIT", sourceAccountTransaction.TransactionType)
		require.Equal(t, account1.ID, sourceAccountTransaction.AccountID)
		require.Equal(t, amount, sourceAccountTransaction.Amount)
		require.NotZero(t, sourceAccountTransaction.ID)

		_, err = store.GetTransaction(context.Background(), sourceAccountTransaction.ID)
		require.NoError(t, err)

		destinationAccountTransaction := result.DestinationTransaction
		require.NotEmpty(t, destinationAccountTransaction)
		require.Equal(t, "CREDIT", destinationAccountTransaction.TransactionType)
		require.Equal(t, account2.ID, destinationAccountTransaction.AccountID)
		require.Equal(t, amount, destinationAccountTransaction.Amount)
		require.NotZero(t, destinationAccountTransaction.ID)

		_, err = store.GetTransaction(context.Background(), destinationAccountTransaction.ID)
		require.NoError(t, err)

		//go:TODO check account balance
	}

}