package db

import (
	"context"
	"database/sql"
)

/*
	The idea here is to create a repository that uses a db instance or transaction instance to repository the database
*/
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository struct {
	db DB
}

func CreateNew(db DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateWithTransaction(dbTxn *sql.Tx) *Repository {
	return &Repository{db: dbTxn}
}

