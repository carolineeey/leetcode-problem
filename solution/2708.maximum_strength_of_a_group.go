package solution

import "sort"

func MaxStrength(nums []int) int64 {
	negatives := []int{}
	strength := int64(1)
	hasPositive := false
	for _, num := range nums {
		if num > 0 {
			strength *= int64(num)
			hasPositive = true
		} else if num < 0 {
			negatives = append(negatives, num)
		}
	}

	// sort negatives to pair the largest (least negative) ones
	sort.Slice(negatives, func(i, j int) bool {
		return negatives[i] < negatives[j]
	})

	// pair negatives to get positive products
	for i := 0; i+1 < len(negatives); i += 2 { // skip two numbers (to avoid duplicates)
		strength *= int64(negatives[i] * negatives[i+1])
	}

	// if no positives and not enough negatives to pair, pick the largest negative
	if !hasPositive && len(negatives) < 2 {
		maxVal := nums[0]
		for _, num := range nums {
			if num > maxVal {
				maxVal = num
			}
		}
		return int64(maxVal)
	}

	return strength
}
