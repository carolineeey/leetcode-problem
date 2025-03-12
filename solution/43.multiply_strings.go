package solution

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// initialize result array as its size will be at most len(num1)+len(num2)
	result := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// convert char digits to int and calculate the product
			product := int(num1[i]-'0') * int(num2[j]-'0')
			sum := product + result[i+j+1]

			// carry over the value to the next position
			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	// convert the result to string and skip leading zeros
	resultStr := ""
	for _, v := range result {
		if !(len(resultStr) == 0 && v == 0) {
			resultStr += string(rune(v + '0'))
		}
	}

	return resultStr
}
