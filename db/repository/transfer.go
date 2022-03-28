package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
)

const createTransfer = `
INSERT INTO transfers ( source_account_id,destination_account_id,amount, currency, reference) 
VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, source_account_id, destination_account_id, amount, currency, created_at
`


type CreateTransferParams struct {
	SourceAccountID int64 `json:"source_account_id"`
	DestinationAccountID   int64 `json:"destination_account_id"`
	Amount        int64 `json:"amount"`
	Currency 	string 	`json:"currency"`
	Reference string `json:"reference"`
}

func (r *Repository) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {


	row := r.db.QueryRowContext(ctx, createTransfer, arg.SourceAccountID, arg.DestinationAccountID, arg.Amount, arg.Currency, arg.Reference)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.SourceAccountID,
		&i.DestinationAccountID,
		&i.Amount,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `
SELECT id, source_account_id, destination_account_id, amount, currency, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (r *Repository) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := r.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.SourceAccountID,
		&i.DestinationAccountID,
		&i.Amount,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `
SELECT id, source_account_id, destination_account_id, amount, currency, created_at FROM transfers
WHERE 
    source_account_id = $1 OR
    destination_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	SourceAccountID int64 `json:"source_account_id"`
	DestinationAccountID   int64 `json:"destination_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (r *Repository) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := r.db.QueryContext(ctx, listTransfers,
		arg.SourceAccountID,
		arg.DestinationAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.SourceAccountID,
			&i.DestinationAccountID,
			&i.Amount,
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