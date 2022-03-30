package db

import (
	"context"
	. "github.com/Adetunjii/simplebank/db/models"
)

const createUser = `
	INSERT INTO users (username, password, email, full_name) 
	VALUES ($1, $2, $3, $4) 
    RETURNING *
`

type CreateUserDto struct {
	Username    string   	`json:"username"`
	Password  	string    	`json:"password"`
	Email 	 	string   	`json:"email"`
	FullName 	string 		`json:"full_name"`
}

func (r *Repository) CreateUser(ctx context.Context, arg CreateUserDto) (User, error) {

	var user User
	err := r.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email, arg.FullName).Scan(
		&user.Username,
		&user.Password,
		&user.FullName,
		&user.Email,
		&user.CreatedAt,
	)
	return user, err
}

const getUser = `
	SELECT username, password, email, full_name, created_at  FROM users
	WHERE username = $1 LIMIT 1
`
func (r *Repository) GetUser(ctx context.Context, username string) (User, error) {
	row := r.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.Email,
		&i.FullName,
		&i.CreatedAt,
	)

	return i, err
}
