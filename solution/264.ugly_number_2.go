package solution

func NthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1 // first ugly number is 1

	i2, i3, i5 := 0, 0, 0 // pointers for 2,3,5

	for i := 1; i < n; i++ {
		// candidates for each pointer
		next2 := ugly[i2] * 2
		next3 := ugly[i3] * 3
		next5 := ugly[i5] * 5

		// pick the smallest of ugly number
		nextUgly := min(next2, min(next3, next5))
		ugly[i] = nextUgly

		// move any pointer that produce this value
		if nextUgly == next2 {
			i2++
		}

		if nextUgly == next3 {
			i3++
		}

		if nextUgly == next5 {
			i5++
		}
	}

	return ugly[n-1]
}
