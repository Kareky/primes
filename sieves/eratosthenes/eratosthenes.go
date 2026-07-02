package eratosthenes

import "github.com/Kareky/primes/internal/errors"

//FindPrimes returns all prime numbers up to upperBound.
//It returns error if upperBound exceed 1,000,000,000.
func FindPrimes(upperBound int) ([]int, error) {
	if upperBound > SizeLimit {
		return nil, errors.ErrMaxSizeExceed(packageName, SizeLimit)
	}

	if upperBound < 2 {
    	return []int{}, nil
	}

	var primeList = []int{}
	primeList = append(primeList, 2)
	// Slice stores odd numbers starting at 3: index i represents 2*i+3
    // true = marked composite, false = still prime candidate
	var numberList = make([]bool, (upperBound-1)/2)
	
	for i := 0; i < len(numberList); i++ {
		if numberList[i] {
			continue // composite number, aka true value indexes, are skipped
		}

		p := i*2+3
		primeList = append(primeList, p)
	
		// Only need to mark composites for primes up to sqrt(upperBound)
		if p*p <= upperBound {
			// Start marking at p*p (smaller multiples already marked by smaller primes).
			// Skip even numbers by stepping 2*p.
			for c := p*p; c <= upperBound; c+=2*p {
				j := (c-3) /2
				numberList[j] = true
			}
		}
	}

	return primeList, nil
}