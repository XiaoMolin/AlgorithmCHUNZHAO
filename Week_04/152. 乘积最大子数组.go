package main

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

func main() {

}

// 同最大子序列和
// 不过这里分两种情况
// 最小负数*负数
// 最大正数*正数
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxv, minv, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmp1 := max(nums[i]*maxv, max(nums[i]*minv, nums[i]))
		tmp2 := min(nums[i]*maxv, min(nums[i]*minv, nums[i]))
		maxv = tmp1
		minv = tmp2
		res = max(maxv, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
