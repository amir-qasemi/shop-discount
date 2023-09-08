package db

import (
	"database/sql"

	"github.com/amir-qasemi/shop-discount/internal/config"
)

// Connect to the database according to the given config.
// Not implemented yet
func Connect(config config.DbConfig) (*sql.DB, error) {
	return nil, nil
}
