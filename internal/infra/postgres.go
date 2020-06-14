package infra

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("POSTGRES_CONNECTION")
	if connStr == "" {
		return nil, errors.New("please provide a db connection string")
	}

	// postgres://postgres:123456@localhost/postgres?sslmode=disable

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, err
}
