package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewInmemoryDb() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("cannot connect to sqlite3, %w", err)
	}
	return db
}
