package solution

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		index := -1
		for a, v := range nums2 {
			if v == num {
				index = a
				break
			}
		}

		// search for next greater element
		next := -1
		for j := index + 1; j < len(nums2); j++ {
			if nums2[j] > num {
				next = nums2[j]
				break
			}
		}

		res[i] = next
	}

	return res
}
