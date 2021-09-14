package repository

import (
	"database/sql"
)

type BankPostgres struct {
	db *sql.DB
}

func NewRoomPostgres(db *sql.DB) *BankPostgres {
	return &BankPostgres{db: db}
}
