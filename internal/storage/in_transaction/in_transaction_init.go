package in_transaction

import (
	"database/sql"
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return Transaction{
		db: db,
	}
}
