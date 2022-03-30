package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
)

type IRepository interface {
	CreateAccount(ctx context.Context, arg CreateAccountDto) (Account, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	ListAccounts(ctx context.Context, arg ListAccountParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountDto) (Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	CreateTransaction(ctx context.Context, arg CreateTransactionDto) (Transaction, error)
	GetTransaction(ctx context.Context, id int64) (Transaction, error)
	ListTransactions(ctx context.Context, arg ListTransactionParams) ([]Transaction, error)
	UpdateTransaction(ctx context.Context, arg UpdateTransactionDto) (Transaction, error)
	DeleteTransaction(ctx context.Context, id int64) error
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserDto) (User, error)
}