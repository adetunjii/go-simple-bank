package db

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/Adetunjii/simplebank/db/models"
	"github.com/Adetunjii/simplebank/util"
)

// An extension to the repository where we can run all queries as well as transactions
type Store struct {
	*Repository			//embed repository
	db *sql.DB
}

func CreateNewStore(db *sql.DB) *Store {
	return &Store {
		db: db,
		Repository: CreateNew(db),
	}
}


func (store *Store) execTxn(ctx context.Context, fn func(*Repository) error) error {
	txn, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	repository := CreateNew(txn)
	err = fn(repository)

	if err != nil {
		if rbErr := txn.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return txn.Commit()
}


type TransferTxnParams struct {
	SourceAccountID int64 	`json:"source_account_id"`
	DestinationAccountID int64 	`json:"destination_account_id"`
	Amount int64 		`json:"amount"`
	Currency string		`json:"currency"`
	Reference string	`json:"reference"`
}

type TransferTxnResult struct {
	Transfer Transfer		 				`json:"transfer"`
	SourceAccount Account					`json:"source_account"`
	DestinationAccount Account 				`json:"destination_account"`
	SourceTransaction  Transaction			`json:"source_transaction"`
	DestinationTransaction Transaction 		`json:"destination_transaction"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxnParams) (TransferTxnResult, error) {
	var result TransferTxnResult

	err := store.execTxn(ctx, func(r *Repository) error {

		var err error

		result.Transfer, err = r.CreateTransfer(ctx, CreateTransferParams{
			SourceAccountID: arg.SourceAccountID,
			DestinationAccountID: arg.DestinationAccountID,
			Amount: arg.Amount,
			Currency: arg.Currency,
			Reference: arg.Reference,
		})

		if err != nil {
			return err
		}

		result.SourceTransaction, err = r.CreateTransaction(ctx, CreateTransactionDto{
			AccountID:       arg.SourceAccountID,
			Amount:          arg.Amount,
			TransactionType: "DEBIT",
			Currency:        arg.Currency,
			Status:          "PENDING",
			Reference:       arg.Reference,
		})

		if err != nil {
			return err
		}

		result.DestinationTransaction, err = r.CreateTransaction(ctx, CreateTransactionDto{
			AccountID:       arg.DestinationAccountID,
			Amount:          arg.Amount,
			TransactionType: "CREDIT",
			Currency:        arg.Currency,
			Status:          "PENDING",
			Reference:       util.RandomString(10),
		})

		if err != nil {
			return err
		}

		// go:TODO update account balance and transactions to successful

		return nil
	})

	return result, err
}