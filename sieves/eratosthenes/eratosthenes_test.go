package eratosthenes_test

import (
	"testing"

	era "github.com/Kareky/primes/sieves/eratosthenes"
)

func TestFindPrimes(t *testing.T) {
	tests := []struct {
		name     string
		bound    int
		want     []int
		wantErr  bool
	}{
		{
			name:     "less than 2",
			bound:    1,
			want:     []int{},
			wantErr:  false,
		},
		{
			name:     "bound 2",
			bound:    2,
			want:     []int{2},
			wantErr:  false,
		},
		{
			name:     "bound 3",
			bound:    3,
			want:     []int{2, 3},
			wantErr:  false,
		},
		{
			name:     "bound 10",
			bound:    10,
			want:     []int{2, 3, 5, 7},
			wantErr:  false,
		},
		{
			name:     "bound 30",
			bound:    30,
			want:     []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
			wantErr:  false,
		},
		{
			name:     "bound 100",
			bound:    100,
			want:     []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			wantErr:  false,
		},
		{
			name:     "exceeds size limit",
			bound:    era.SizeLimit,
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := era.FindPrimes(tt.bound)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindPrimes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("FindPrimes() length = %d, want %d", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("FindPrimes() at index %d = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFindPrimes_Completeness(t *testing.T) {
	// Ensures all numbers marked as prime are actually prime
	// and all primes ≤ bound are present
	primes, err := era.FindPrimes(1000)
	if err != nil {
		t.Fatalf("FindPrimes(1000) failed: %v", err)
	}

	// Check sorted order
	for i := 1; i < len(primes); i++ {
		if primes[i] <= primes[i-1] {
			t.Errorf("primes not sorted: %d > %d at indices %d, %d", primes[i-1], primes[i], i-1, i)
		}
	}

	// Check each prime is actually prime (simple trial division)
	for _, p := range primes {
		for d := 2; d*d <= p; d++ {
			if p%d == 0 {
				t.Errorf("%d is not prime (divisible by %d)", p, d)
			}
		}
	}

	// Check no primes are missing: walk up to 1000 and verify
	// every prime is in the list
	primeMap := make(map[int]bool)
	for _, p := range primes {
		primeMap[p] = true
	}

	for n := 2; n <= 1000; n++ {
		isPrime := true
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime && !primeMap[n] {
			t.Errorf("%d is prime but missing from result", n)
		}
		if !isPrime && primeMap[n] {
			t.Errorf("%d is composite but included in result", n)
		}
	}
}

func TestFindPrimes_ErrorReturned(t *testing.T) {
	_, err := era.FindPrimes(era.SizeLimit+1)
	if err == nil {
		t.Errorf("FindPrimes(%d) expected error, got nil", era.SizeLimit+1)
	}
}

func BenchmarkFindPrimes_1e7(b *testing.B) {
    for b.Loop() {
        era.FindPrimes(10000000)
    }
}

func BenchmarkFindPrimes_1e9(b *testing.B) {
    for b.Loop() {
        era.FindPrimes(1000000000)
    }
}