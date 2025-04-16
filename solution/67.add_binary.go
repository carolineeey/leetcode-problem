package solution

func AddBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	var result []byte

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0') // convert '0' or '1' char to int
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// append the current digit
		result = append(result, byte(sum%2+'0'))

		// update carry
		carry = sum / 2
	}

	// reverse the result
	for i, n := 0, len(result); i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}

	return string(result)
}
