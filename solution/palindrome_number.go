package solution

func IsPalindrome(number int) (isPalindrome bool) {
	remainder := 0
	temp := 0
	reverse := 0
	temp = number
	for {
		if number <= 0 {
			break
		}

		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10
	}

	if temp == reverse {
		return true
	} else {
		return false
	}
}
