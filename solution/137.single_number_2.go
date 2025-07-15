package solution

func SingleNumber2(nums []int) int {
	var ones, twos int

	for _, num := range nums {
		ones = (ones ^ num) &^ twos // add num to ones if it's not in twos
		twos = (twos ^ num) &^ ones // add num to twos if it's not in ones
	}

	return ones
}
