package db

import (
	"database/sql"
)

// DB represents a database connection and provides methods to interact with the database.
type DB struct {
	db *sql.DB
}

// Default is the default database connection used throughout the application.
var Default *DB