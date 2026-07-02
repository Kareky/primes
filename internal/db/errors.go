package db

import (
	"errors"
)

var (
	ErrNoRowsAffected = errors.New("no rows affected by query")
)