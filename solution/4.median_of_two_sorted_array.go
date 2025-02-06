package solution

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge both array into one
	sortedArrays := append(nums1, nums2...)
	sort.Ints(sortedArrays)

	// get the length of the merged array
	n := len(sortedArrays)
	mid := n / 2

	// if odd, return the middle element
	if n%2 == 1 {
		return float64(sortedArrays[mid])
	}

	// if even, return the average of the two middle elements
	return float64(sortedArrays[mid-1]+sortedArrays[mid]) / 2.0
}
