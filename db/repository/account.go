package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
)

// CREATE AN ACCOUNT
const createAccount = `
	INSERT INTO accounts (owner, balance, currency) 
	VALUES ($1, $2, $3) 
    RETURNING id, owner, balance, currency, created_at
`

type CreateAccountDto struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (r *Repository) CreateAccount(ctx context.Context, arg CreateAccountDto) (Account, error) {

	var account Account
	err := r.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency).Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err
}

// FETCH AN ACCOUNT BY ID
const getAccount = `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
`

func (r *Repository) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := r.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)

	return i, err
}

const getAccountForUpdate = `
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE`

func (r *Repository) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := r.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE owner = $1
	ORDER BY id
	LIMIT $2
	OFFSET $3
`

type ListAccountParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (r *Repository) ListAccounts(ctx context.Context, arg ListAccountParams) ([]Account, error) {
	rows, err := r.db.QueryContext(ctx, listAccounts, arg.Owner, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Account{}

	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
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

const updateAccount = `
	UPDATE accounts
	SET balance = $2
	WHERE ID = $1
	RETURNING id, owner, balance, currency, created_at
`

type UpdateAccountDto struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (r *Repository) UpdateAccount(ctx context.Context, arg UpdateAccountDto) (Account, error) {
	row := r.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `
	DELETE FROM accounts WHERE id = $1
`

func (r *Repository) DeleteAccount(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, deleteAccount, id)
	return err
}
