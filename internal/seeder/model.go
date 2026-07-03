package seeder

import (
	"github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/sieves/eratosthenes"
)

// Seed represents a seeder that can seed the database with prime numbers.
type Seed struct {
	db *db.DB
}

// NewSeed creates a new Seed instance with the provided database connection.
func NewSeed(db *db.DB) Seed {
	return Seed{db: db}
}

// Handler is a function type that defines the signature for prime number generation algorithms.
type Handler func(upperBound int) ([]int, error)

// Handlers is a map that associates algorithm names with their corresponding prime number generation functions.
var Handlers = map[string]Handler{
	"eratosthenes": eratosthenes.FindPrimes,
}