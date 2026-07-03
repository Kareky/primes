package division

import "github.com/Kareky/primes/internal/db"

const algorithm = "division primalty test"
var sizeLimit int

// UpdateSizeLimit updates the size limit for the division primality test based on the maximum prime number stored in the database.
func UpdateSizeLimit() error {
	maxPrime, err := db.Default.GetMaxPrime()
	if err != nil {
		return err
	}
	sizeLimit = maxPrime * maxPrime
	return nil
}