package main

import "fmt"

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//示例 2：
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//提示：
//
//1 <= nums.length <= 10
//-10 <= nums[i] <= 10
//nums 中的所有元素 互不相同

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs([]int{}, nums, 0, &res)
	return res
}

func dfs(data []int, nums []int, index int, res *[][]int) {
	//必须复制一份
	temp := make([]int, len(data))
	copy(temp, data)
	*res = append(*res, temp)
	for i := index; i < len(nums); i++ {
		data = append(data, nums[i])
		dfs(data, nums, i+1, res)
		data = data[:len(data)-1] //回溯
	}
}
