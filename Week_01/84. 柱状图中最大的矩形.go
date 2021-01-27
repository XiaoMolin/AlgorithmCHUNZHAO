package main

func main() {
	largestRectangleArea([]int{1, 3, 4, 56, 6, 2})
}

func largestRectangleArea(heights []int) int {
	max := 0
	stack := []int{0}
	heights = append([]int{-1}, heights...)
	heights = append(heights, 0)
	for i := 1; i < len(heights); i++ {
		if heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for {

			if heights[stack[len(stack)-1]] < heights[i] {
				stack = append(stack, i)
				break
			}

			top := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			area := (i - stack[len(stack)-1] - 1) * top
			if area > max {
				max = area
			}

		}
	}

	return max
}
