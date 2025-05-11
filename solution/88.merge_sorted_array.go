package solution

func MergeSorted(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	k := m + n - 1 // last index

	// merge in reverse order
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i] // go to last position and replace with the bigger one
			i--                 // decrement as we merge in reverse order
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// if any elements left in nums2, fill nums1 with leftover nums2 element
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
