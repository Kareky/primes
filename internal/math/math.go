package math

// ModularExp calculates (base^exp) mod mod using the method of exponentiation by squaring.
// It takes three integers as input: base, mod, and exp, and returns the result of the modular exponentiation.
func ModularExp(base, mod, exp int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp % 2 == 1 {
			result = (base * result) % mod
		}

		base = (base * base) % mod
		exp = exp / 2
	}

	return result
}

// GDC returns the greatest common divisor of two integers using the Euclidean algorithm.
// It takes two integers n1 and n2 as input and returns their GCD.
func GCD(n1, n2 int) int {
	for n2 != 0 {
        n1, n2 = n2, n1%n2
    }
    return n1
}