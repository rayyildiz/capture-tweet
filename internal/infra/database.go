package infra

import (
	"capturetweet.com/internal/ent"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewPostgresDatabase() *ent.Client {
	dsn := os.Getenv("POSTGRES_CONNECTİON")
	db, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("can not open connection, %v", err)
	}
	return db
}
