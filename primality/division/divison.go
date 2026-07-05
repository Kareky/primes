package division

import (
	"log"
	"math"

	"github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/internal/errors"
)

// DBIsPrime checks if a number is prime using the division method. It first checks if the number exceeds the size limit,
// then checks if it exists in the database of known primes.
// If not, it retrieves all primes up to the square root of the number and checks for divisibility.
// Returns true if the number is prime, false otherwise. If an error occurs during database operations, it returns the error.
// This version is faster than normal trial divion but occupy more space.
// For very large inputs, which are memory intensive, use 'IsPrime' instead
func IsPrimeDB(number int) (bool, error) {
	if number > sizeLimit {
		return false, errors.ErrMaxSizeExceed(algorithm, sizeLimit)
	}

	if number < 2 {
		return false, nil
	}

	exists, err := db.Default.Exists(number)
	if err != nil {
		return false, err
	}

	if exists {
		return true, nil
	}

	if primeList == nil {
		err = CachePrimes()
		if err != nil {
			return false, err
		}
	}

	log.Println("Finished getting all primes from db")

	limit := int(math.Sqrt(float64(number)))

	for _, prime := range primeList {
		if prime > limit {
			break
		}

		if number % prime == 0 {
			return false, nil
		}
	}

	log.Println("Finished loop")

	return true, nil
}

// IsPrime checks if a number is prime using the trial division method.
// It returns true if the number is prime, false otherwise.
// This method is slower than 'DBIsPrime' but uses no permanent memory, making it suitable for very large inputs.
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}

	if number % 2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i += 2 {
		if number % i == 0 {
			return false
		}
	}

	return true
}