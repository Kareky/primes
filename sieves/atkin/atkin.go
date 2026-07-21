package atkin

import "github.com/Kareky/primes/internal/errors"

// FindPrimes returns all prime numbers up to upperBound using the Sieve of Atkin algorithm.
// It returns an error if upperBound exceeds 1,000,000,000.
func FindPrimes(upperBound int) ([]int, error) {
	if upperBound > SizeLimit {
		return nil, errors.ErrMaxSizeExceed(packageName, SizeLimit)
	}

	if upperBound < 2 {
		return []int{}, nil
	}

	var primeList = []int{}
	if upperBound == 2 {
		primeList = append(primeList, 2)
		return primeList, nil
	}

	primeList = append(primeList, 2, 3)
	if upperBound == 3 {
		return primeList, nil
	}

	// Slice stores odd numbers starting at 3: index i represents 2*i+3
    // true = marked composite, false = still prime candidate
	var numberList = make([]bool, upperBound+1)

	for x := 1; x * x <= upperBound; x++ {
		for y := 1; y * y <= upperBound; y++ {
			n := (4 * x * x) + (y * y)
			if n <= upperBound && (n % 12 == 1 || n % 12 == 5) {
				numberList[n] = !numberList[n]
			}

			n = (3 * x * x) + (y * y)
			if n <= upperBound && n % 12 == 7 {
				numberList[n] = !numberList[n]
			}

			n = (3 * x * x) - (y * y)
			if n <= upperBound && x > y && n % 12 == 11 {
				numberList[n] = !numberList[n]
			}
		}
	}

	for n := 5; n <= upperBound; n+=2 {
		if numberList[n] {
			if n * n <= upperBound {
				for i := n * n; i <= upperBound; i += n * n {
					numberList[i] = false
				}
			}

			primeList = append(primeList, n)
		}
	}

	return primeList, nil
}