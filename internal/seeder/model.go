package seeder

import (
	"github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/sieves/eratosthenes"
)

type Seed struct {
	db *db.DB
}

func NewSeed(db *db.DB) Seed {
	return Seed{db: db}
}

type Handler func(upperBound int) ([]int, error)

var Handlers = map[string]Handler{
	"eratosthenes": eratosthenes.FindPrimes,
}