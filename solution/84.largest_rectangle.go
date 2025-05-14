package solution

func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{}
	heights = append(heights, 0) // add a 0-height bar to flush the stack at the end

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] { // look for smaller height
			topIndex := stack[len(stack)-1] // topIndex is the index of the bar we are now done with, it's the top of the stack
			stack = stack[:len(stack)-1]    // pop topIndex from the stack (remove last element)

			height := heights[topIndex] // get the height of the bar we just popped
			width := i
			if len(stack) > 0 {
				// stack[len(stack)-1] is the previous index still in the stack
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		if i < len(heights) {
			stack = append(stack, i)
		}
	}

	return maxArea
}
