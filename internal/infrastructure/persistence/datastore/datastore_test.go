package datastore

import (
	"go-example/internal/platform/database"

	"github.com/jmoiron/sqlx"
)

func setupDB() (*sqlx.DB, error) {
	return database.Connect(database.ReadConfig())
}
