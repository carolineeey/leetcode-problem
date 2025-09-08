package solution

func CountPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true // the assumption is: at the start, every number is prime until proven otherwise.
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i { // cross out multiplied number as its of course not a prime number
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
