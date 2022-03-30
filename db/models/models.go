package db

import "time"


type TransactionType string
type Currency string

const (
	Credit TransactionType = "CREDIT"
	Debit			       = "DEBIT"
)

const (
	NGN Currency = "NGN"
	USD		     = "USD"
	GHC			 = "GHC"
)

type Account struct {
	ID int64			`json:"id"`
	Owner string		`json:"owner"`
	Balance int64 		`json:"balance"`
	Currency string  	`json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	ID int64						`json:"id"`
	AccountID int64 				`json:"account_id"`
	Amount int64 					`json:"amount"`
	TransactionType string 			`json:"transaction_type"`
	Currency string					`json:"currency"`
	Status string 					`json:"status"`
	Reference string 				 `json:"reference"`
	CreatedAt time.Time 			`json:"created_at"`
}

type Transfer struct {
	ID int64					`json:"id"`
	SourceAccountID int64 		`json:"source_account_id"`
	DestinationAccountID int64 	`json:"destination_account_id"`
	Amount int64				`json:"amount"`
	Currency string				`json:"currency"`
	Reference string			`json:"reference"`
	CreatedAt time.Time 		`json:"created_at"`
}

type User struct {
	Username    string   	`json:"username"`
	Password  	string    	`json:"balance"`
	FullName 	string 		`json:"full_name"`
	Email 	 	string   	`json:"currency"`
	CreatedAt   time.Time 	`json:"created_at"`
}