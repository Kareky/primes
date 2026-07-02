package seeder

import (
    "github.com/schollz/progressbar/v3"
	"log"
)

// PopulatePrimesSeed populates the database with prime numbers up to the specified upperBound using the specified algorithm.
func (s Seed) PopulatePrimesSeed(upperBound int, algorithm string) {
	handler, exists := Handlers[algorithm]
	if !exists {
		log.Fatal("Algorithm not found: " + algorithm)
	}

	log.Printf("Seeding with %s...", algorithm)
	primes, err := handler(upperBound)
	if err != nil {
		log.Fatal("Error during primes generation: " + err.Error())
	}

    bar := progressbar.Default(int64(len(primes)))
    err = s.db.InsertPrimes(primes, func(current, total int) {
        bar.Set(current)
    })
	if err != nil {
		log.Fatal("Error during save: " + err.Error())
	}

	log.Printf("Seed with %s succeed", algorithm)
}