package solution

func AddToArrayForm(num []int, k int) []int {
	i := len(num) - 1
	res := []int{}

	for i >= 0 || k > 0 {
		if i >= 0 {
			k += num[i]
			i--
		}
		res = append(res, k%10)
		k /= 10
	}

	// reverse the result since we added digits from least to most significant
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return res
}
