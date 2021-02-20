package main

//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
//示例 1：
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

func main() {

}

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
