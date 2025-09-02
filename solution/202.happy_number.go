package solution

func IsHappy(n int) bool {
	slow := n
	fast := getNext(n)

	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

func getNext(n int) (next int) {
	total := 0
	for n > 0 {
		d := n % 10
		total += d * d
		n /= 10
	}

	return total
}
