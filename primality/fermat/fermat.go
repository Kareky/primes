package fermat

import (
	"log"
	"math/rand/v2"
	"github.com/Kareky/primes/internal/math"
)

// IsPrime checks if a number is prime using the Fermat primality test.
// It first checks if the number is less than 2, in which case it returns false.
// It then performs a specified number of repetitions (default is 5) of the test,
// where it randomly selects a base up to (number-1) and checks if the modular exponentiation condition holds.
// If any repetition fails, it returns false, indicating that the number is composite.
// If all repetitions pass, it returns true, indicating that the number is likely prime.
func IsPrime(number int, iter ...int) bool {
	if number < 2 {
		return false
	}

	reps := 5
    if len(iter) > 0 && iter[0] > 0 {
        reps = iter[0]
    }

	//a^(number-1) mod p = 1
	for i := range reps {
		base := 2 + rand.IntN(number-3)

		if math.GCD(base, number) != 1 {
			return false
		}
		
		if math.ModularExp(base, number, number-1) != 1 {
			return false
		}
		log.Printf("Repetition number %d", i)
	}

	return true
}