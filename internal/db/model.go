package db

import (
	"database/sql"
)

type DB struct {
	db *sql.DB
}

var Default *DB