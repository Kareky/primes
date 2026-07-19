package millerrabin

import (
	"math/rand/v2"

	"github.com/Kareky/primes/internal/math"
)

// IsPrime checks if a number is prime using the Miller-Rabin primality test.
// It first checks if the number is less than 2, in which case it returns false.
// If the number is even and greater than 2, it returns false.
// If the number is 2 or 3, it returns true.
// It then performs a specified number of repetitions (default is 5) of the test,
// where it randomly selects a base and checks if the Miller-Rabin conditions hold.
// If any repetition fails, it returns false, indicating that the number is composite.
// If all repetitions pass, it returns true, indicating that the number is likely prime.
func IsPrime(number int, iter ...int) bool {
	if number <= 3 {
		return number > 1
	}

	if number % 2 == 0 {
		return false
	}

	reps := 5
	if len(iter) > 0 && iter[0] > 0 {
		reps = iter[0]
	}

	var d = number - 1
	var s = 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	for i := 0; i < reps; i++ {
		if !millerRabin(number, d, s) {
			return false
		}
	}

	return true
}

// millerRabin performs the Miller-Rabin primality test for a given number and a value d.
// It randomly selects a base and checks if the Miller-Rabin conditions hold.
// If the conditions are met, it returns true, indicating that the number is likely prime.
// If the conditions are not met, it returns false, indicating that the number is composite.
func millerRabin(number, d, s int) bool {
	a := 2 + rand.IntN(number-3)

	x := math.ModularExp(a, number, d)
	if x == 1 || x == number-1 {
		return true
	}

	for r := 1; r < s; r++ {
		x = (x * x) % number
		if x == number-1 {
			return true
		}
		if x == 1 {
			return false
		}
	}
	return false
}