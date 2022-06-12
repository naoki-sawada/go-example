package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ReadConfig() string {
	return os.Getenv("DATABASE_URL")
}

func Connect(dsn string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dsn)
}
