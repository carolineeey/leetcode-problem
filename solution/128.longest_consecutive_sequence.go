package solution

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		// checks if the current number num is the start of a new consecutive sequence
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// increase currentNum as long as the next consecutive number exists in the set
			for numSet[currentNum+1] {
				currentStreak++
				currentNum++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}

func LongestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	count := 1
	maxCount := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		} else if nums[i]+1 == nums[i+1] {
			count++
		} else {
			if maxCount < count {
				maxCount = count
			}

			count = 1
		}
	}

	if maxCount < count {
		maxCount = count
	}

	return maxCount
}
