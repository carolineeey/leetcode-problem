package solution

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		minHeight := 0
		width := right - left
		if height[left] < height[right] {
			minHeight = height[left]
			left++
		} else {
			minHeight = height[right]
			right--
		}

		area := minHeight * width

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
