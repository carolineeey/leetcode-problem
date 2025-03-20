package solution

func MyPow(x float64, n int) float64 {
	if n == 0 { // base case
		return 1.0
	}

	if n < 0 { // handle negative exponent
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n%2 != 0 { // if n is odd
			result *= x
		}

		x *= x // square the base
		n /= 2 // reduce n by half in each iteration
	}

	return result
}
