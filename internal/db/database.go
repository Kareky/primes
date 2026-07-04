// A simple package for a database of prime numbers, used by algorithms inside the libraries
package db

import (
	"database/sql"
	"fmt"
	"math"
	"strings"

	"github.com/Kareky/primes/config"
)

// DB represents a database connection and provides methods to interact with the database.
// NewDB creates a new instance of DB and initializes the database by creating the necessary tables.
func NewDB(db *sql.DB) (*DB, error) {
	newDB := &DB{db: db}
	err := newDB.CreateTable()
	if err != nil {
		return nil, err
	}

	return newDB, nil
}

// Initialize initializes the database connection and sets the Default variable to the new DB instance.
func Initialize(databasePath string) error {
	if databasePath == "" {
		databasePath = config.Config.Database.Path
	}

	sqlDB, err := sql.Open(config.Config.Database.Type, databasePath)
	if err != nil {
		return err
	}

	stmts := []string{
		"PRAGMA temp_store = MEMORY;",
		"PRAGMA mmap_size = 268435456;",
		"PRAGMA cache_size = -64000;",
		"PRAGMA synchronous = OFF;",
		"PRAGMA journal_mode = WAL;",
	}

	for _, s := range stmts {
		sqlDB.Exec(s)
	}

	Default, err = NewDB(sqlDB)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the database connection if it is open.
func Close() error {
	if Default != nil && Default.db != nil {
		return Default.db.Close()
	}
	return nil
}

// CreateTable creates the "primes" table in the database if it does not already exist.
func (d *DB) CreateTable() error {
	_, err := d.db.Exec(`CREATE TABLE IF NOT EXISTS primes (number INTEGER PRIMARY KEY)`)
	return err
}

// InsertPrime inserts a single prime number into the "primes" table.
func (d *DB) InsertPrime(number int) error {
	_, err := d.db.Exec(`INSERT INTO primes (number) VALUES (?)`, number)
	return err
}

// InsertPrimes inserts multiple prime numbers into the "primes" table.
func (d *DB) InsertPrimes(numbers []int, onProgress func(int, int)	) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT or IGNORE INTO primes (number) VALUES (?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i, number := range numbers {
		_, err := stmt.Exec(number)
		if err != nil {
			tx.Rollback()
			return err
		}

		if onProgress != nil {
            onProgress(i+1, len(numbers)) // call after each insert
        }
	}

	return tx.Commit()
}

// Check if a number exists inside the database
func (d *DB) Exists(number int) (bool, error) {
	query := `SELECT EXISTS(
				SELECT 1
				FROM primes
				WHERE number = ?
				)`
	var exists bool
	err := d.db.QueryRow(query, number).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// GetMaxPrime retrieves the maximum prime number stored in the "primes" table.
func (d *DB) GetMaxPrime() (int, error) {
	query := `	SELECT MAX(number)
				FROM primes`
	var maxPrime int
	err := d.db.QueryRow(query).Scan(&maxPrime)
	if err != nil {
		return 0, err
	}

	return maxPrime, nil
}

// GetAllPrimes retrieves all prime numbers from the "primes" table.
func (d *DB) GetAllPrimes() ([]int, error) {
	query := 	`SELECT number
				FROM primes`
	return d.getPrimes(query)
}

// GetPrimesUpTo retrieves all prime numbers from the "primes" table that are less than or equal to the specified number.
func (d *DB) GetPrimesUpTo(number int) ([]int, error) {
	query := 	`SELECT number
				FROM primes
				WHERE number <= ?`
	return d.getPrimes(query, number)
}

// getPrimes is a helper function that executes the provided query with optional conditions and returns the resulting prime numbers.
func (d *DB) getPrimes(query string, conditions ...any) ([]int, error) {
	var sliceSize int
    if len(conditions) > 0 {
        if bound, ok := conditions[0].(int); ok {
            sliceSize = int(float64(bound) / math.Log(float64(bound)))
        }
    } else {
		countQuery := strings.Replace(query, "SELECT number", "SELECT COUNT(*)", 1)
    	err := d.db.QueryRow(countQuery, conditions...).Scan(&sliceSize)
		if err != nil {
			return nil, err
		}
	}

	rows, err := d.db.Query(query, conditions...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var primes = make([]int, 0, sliceSize)
	for rows.Next() {
		var number int64
		err = rows.Scan(&number)
		if err != nil {
			return nil, err
		}

		primes = append(primes, int(number))
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return primes, nil
}
