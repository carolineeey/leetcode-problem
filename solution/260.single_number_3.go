package solution

// xor property: x ^ x = 0, and x ^ 0 = x
// example :
// 1 ^ 2 ^ 1 ^ 3 ^ 2 ^ 5
//= (1 ^ 1) ^ (2 ^ 2) ^ 3 ^ 5
//= 0 ^ 0 ^ 3 ^ 5
//= 3 ^ 5

func SingleNumber3(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num // XOR (exclusive or) cancels out duplicates
	}

	diff := xor & -xor

	a, b := 0, 0
	for _, num := range nums {
		if num&diff == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
