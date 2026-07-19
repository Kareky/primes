package fermat_test

import (
	"testing"

	"github.com/Kareky/primes/primality/fermat"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		iter     []int
		expected bool
	}{
		{"negative", -5, nil, false},
		{"zero", 0, nil, false},
		{"one", 1, nil, false},
		{"two", 2, nil, true},
		{"three", 3, nil, true},
		{"prime_5", 5, nil, true},
		{"prime_7", 7, nil, true},
		{"prime_11", 11, nil, true},
		{"prime_13", 13, nil, true},
		{"prime_17", 17, nil, true},
		{"prime_19", 19, nil, true},
		{"prime_23", 23, nil, true},
		{"prime_29", 29, nil, true},
		{"prime_31", 31, nil, true},
		{"composite_4", 4, nil, false},
		{"composite_6", 6, nil, false},
		{"composite_8", 8, nil, false},
		{"composite_9", 9, nil, false},
		{"composite_10", 10, nil, false},
		{"composite_14", 14, nil, false},
		{"composite_15", 15, nil, false},
		{"composite_21", 21, nil, false},
		{"composite_25", 25, nil, false},
		{"composite_27", 27, nil, false},
		{"prime_101", 101, nil, true},
		{"prime_103", 103, nil, true},
		{"prime_107", 107, nil, true},
		{"prime_999983", 999983, nil, true},
		{"composite_100", 100, nil, false},
		{"composite_999999", 999999, nil, false},
		{"composite_999982", 999982, nil, false},
		// Carmichael numbers require many iterations to detect as composite
		{"carmichael_561", 561, []int{20}, false},
		{"carmichael_1105", 1105, []int{20}, false},
		{"carmichael_1729", 1729, []int{20}, false},
		{"carmichael_2465", 2465, []int{20}, false},
		{"carmichael_2821", 2821, []int{20}, false},
		{"carmichael_6601", 6601, []int{20}, false},
		{"prime_with_1_iter", 17, []int{1}, true},
		{"prime_with_10_iter", 19, []int{10}, true},
		{"composite_with_1_iter", 9, []int{1}, false},
		{"carmichael_with_10_iter", 561, []int{10}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fermat.IsPrime(tt.n, tt.iter...)
			if got != tt.expected {
				t.Errorf("IsPrime(%d, %v) = %v, want %v", tt.n, tt.iter, got, tt.expected)
			}
		})
	}
}

func TestFermatStability(t *testing.T) {
	for i := range 20 {
		if !fermat.IsPrime(101, 10) {
			t.Errorf("prime 101 failed on iteration %d", i)
		}
		if fermat.IsPrime(561, 20) {
			t.Errorf("Carmichael 561 passed on iteration %d", i)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for b.Loop() {
		fermat.IsPrime(999983, 5)
	}
}