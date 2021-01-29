package main

import "fmt"

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func main() {
	fmt.Println(permute([]int{3, 2, 1}))
}

//回溯法
func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, path []int, res *[][]int) {
	if len(nums) == 1 {
		path = append(path, nums[0])
		*res = append(*res, path)
		return
	}
	//循环
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		dfs(nums[1:], append(path, nums[0]), res)
		nums[0], nums[i] = nums[i], nums[0]
	}
}
