// A simple package for a database of prime numbers, used by algorithms inside the libraries
package db

import (
	"database/sql"
	"fmt"

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
	rows, err := d.db.Query(query, conditions...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var primes []int
	for rows.Next() {
		var number int
		err = rows.Scan(&number)
		if err != nil {
			return nil, err
		}

		primes = append(primes, number)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return primes, nil
}
