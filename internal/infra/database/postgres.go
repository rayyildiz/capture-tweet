package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("can not connect to postgres, %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error while ping postgres, %w", err)
	}
	return db, err
}
