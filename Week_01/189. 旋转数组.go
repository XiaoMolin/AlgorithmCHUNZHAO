package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	// k大于nums
	if k > len(nums) {
		k = k % len(nums)
	}
	//原地旋转
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	fmt.Println(nums)

}
