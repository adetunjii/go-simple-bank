package db

import (
	"context"
	"fmt"
	"github.com/Adetunjii/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := CreateNewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> account balance::: ", account1.Balance, account2.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxnResult)

	existed := make(map[int]bool)

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


		sourceAccount := result.SourceAccount
		require.NotEmpty(t, sourceAccount)
		require.Equal(t, account1.ID, sourceAccount.ID)

		destinationAccount := result.DestinationAccount
		require.NotEmpty(t, destinationAccount)
		require.Equal(t, account2.ID, destinationAccount.ID)


		fmt.Println(">> after each txn:::", sourceAccount.Balance, destinationAccount.Balance)
		balance1 := account1.Balance - sourceAccount.Balance
		balance2 := destinationAccount.Balance - account2.Balance
		require.Equal(t, balance1, balance2)
		require.True(t, balance1 > 0)
		require.True(t, balance1 % amount == 0)

		k := int(balance1/amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err := testRepository.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount1)

	updatedAccount2, err := testRepository.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount2)

	fmt.Printf(">> Updated balance 	%v	%v", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance - int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance + int64(n)*amount, updatedAccount2.Balance)
}

func TestStore_TransferTxDeadlock(t *testing.T) {
	store := CreateNewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i++ {

		sourceAccountID := account1.ID
		destinationAccountID := account2.ID

		if i % 2 == 0 {
			sourceAccountID = account2.ID
			destinationAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxnParams{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destinationAccountID,
				Amount:               amount,
				Currency:             util.RandomCurrency(),
				Reference: 			util.RandomString(10),
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++{
		err := <- errs
		require.NoError(t, err)

	}

	updatedAccount1, err := testRepository.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount1)

	updatedAccount2, err := testRepository.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount2)

	fmt.Printf(">> Updated balance 	%v	%v", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}