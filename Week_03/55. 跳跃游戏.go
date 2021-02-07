package main

//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。
//
//
//
//示例 1：
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//示例 2：
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//0 <= nums[i] <= 105

func main() {

}

// 从后往前贪心
func canJump1(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	end := len(nums) - 1
	for i := end - 1; i >= 0; i-- {
		if nums[i]+i >= end {
			end = i
		}
	}
	return end == 0
}

// 动态规划bfs
// 时间复杂度O(N*M)
// 效率较低
// 高效率：类似贪心的dp
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if dp[i] == true {
			for j := i; j <= i+nums[i] && j < len(dp); j++ {
				dp[j] = true
			}
		}
	}
	return dp[len(nums)-1]
}
