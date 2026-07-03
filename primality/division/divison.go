package division

import (
	"math"

	"github.com/Kareky/primes/internal/db"
	"github.com/Kareky/primes/internal/errors"
)

// IsPrime checks if a number is prime using the division method. It first checks if the number exceeds the size limit,
// then checks if it exists in the database of known primes.
// If not, it retrieves all primes up to the square root of the number and checks for divisibility.
// Returns true if the number is prime, false otherwise. If an error occurs during database operations, it returns the error.
func IsPrime(number int) (bool, error) {
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

	primeList, err := db.Default.GetPrimesUpTo(int(math.Sqrt(float64(number))))
	if err != nil {
		return false, err
	}

	for _, prime := range primeList {
		if number % prime == 0 {
			return false, nil
		}
	}

	return true, nil
}