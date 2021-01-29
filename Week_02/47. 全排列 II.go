package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
//[1,2,1],
//[2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10

func main() {
	fmt.Println(permuteUnique([]int{3, 2, 1, 2, 2}))
}

//使用回溯加上map
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, path []int, res *[][]int) {
	if len(nums) == 1 {
		path = append(path, nums[0])
		*res = append(*res, dump(path))
		return
	}
	maps := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		if maps[nums[0]] == false {
			maps[nums[0]] = true
			dfs(nums[1:], append(path, nums[0]), res)
		}
		nums[i], nums[0] = nums[0], nums[i]
	}

}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
