package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
)

const createTransaction = `
	INSERT INTO transactions(account_id, amount, transaction_type, currency, status, reference)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, account_id, amount, transaction_type, currency, status, reference, created_at
`

type CreateTransactionDto struct {
	AccountID int64 		 `json:"account_id"`
	Amount int64			 `json:"amount"`
	TransactionType string	 `json:"transaction_type"`
	Currency string 		 `json:"currency"`
	Status string 			`json:"status"`
	Reference string 		`json:"reference"`
}

func (r *Repository) CreateTransaction(ctx context.Context, arg CreateTransactionDto) (Transaction, error) {

	var transaction Transaction
	err := r.db.QueryRowContext(ctx, createTransaction, arg.AccountID, arg.Amount, arg.TransactionType, arg.Currency, arg.Status, arg.Reference).Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.Amount,
			&transaction.TransactionType,
			&transaction.Currency,
			&transaction.Status,
			&transaction.Reference,
			&transaction.CreatedAt,
		)

	return transaction, err
}

const getTransaction = `
	SELECT id, account_id, amount, transaction_type, currency, status, reference, created_at FROM transactions
	WHERE id = $1 LIMIT 1
`

func (r *Repository) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := r.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.TransactionType,
		&i.Currency,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
	)

	return i, err
}

const listTransactions = `
	SELECT * FROM transactions
	ORDER BY id
	LIMIT $1
	OFFSET $2
`

type ListTransactionParams struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}


func (r *Repository) ListTransactions(ctx context.Context, arg ListTransactionParams) ([]Transaction, error) {
	rows, err := r.db.QueryContext(ctx, listTransactions, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Transaction

	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.TransactionType,
			&i.Currency,
			&i.Status,
			&i.Reference,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateTransaction = `
	UPDATE transactions
	SET status = $2
	WHERE ID = $1
	RETURNING id, account_id, amount, transaction_type, currency, status, reference, created_at;
`

type UpdateTransactionDto struct {
	ID int64 `json:"id"`
	Status string `json:"status"`
}

func (r *Repository) UpdateTransaction(ctx context.Context, arg UpdateTransactionDto) (Transaction, error) {
	row := r.db.QueryRowContext(ctx, updateTransaction, arg.ID, arg.Status)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.TransactionType,
		&i.Currency,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransaction = `
	DELETE FROM Transactions WHERE id = $1
`

func (r *Repository) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, deleteTransaction, id)
	return err
}