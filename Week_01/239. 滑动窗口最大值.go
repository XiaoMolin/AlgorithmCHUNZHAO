package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{12, 3, 2, 4, 1}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(q) != 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if q[0] == i-k {
			q = q[1:]
		}
		if i+1-k >= 0 {
			res[i+1-k] = nums[q[0]]
		}
	}
	return res
}
