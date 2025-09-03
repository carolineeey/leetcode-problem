package solution

func AddDigits(num int) int {
	for num >= 10 {
		num = getNextDigits(num)
	}

	return num
}

func getNextDigits(n int) int {
	total := 0
	for n > 0 {
		d := n % 10
		total += d
		n /= 10
	}
	return total
}
