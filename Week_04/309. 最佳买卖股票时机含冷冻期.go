package main

//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//示例:
//
//输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 买入 又冷静期
		if i >= 2 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
