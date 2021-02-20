package main

//给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//示例:
//
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//说明:
//
//假设你总是可以到达数组的最后一个位置。

func main() {

}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//从前往后
	//dp[i]=dp[i+1]+nums[i]
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		if dp[i] >= 0 {
			for j := 1; j <= nums[i] && j+i < len(nums); j++ {
				if dp[i+j] == -1 {
					dp[j+i] = dp[i] + 1
				} else {
					dp[j+i] = min(dp[j+i], dp[i]+1)
				}
				if j+i == len(nums)-1 {
					return dp[j+i]
				}
			}
		}
	}

	return dp[len(nums)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
