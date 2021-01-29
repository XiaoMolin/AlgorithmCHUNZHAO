package main

import "fmt"

//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
//示例:
//
//输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//说明:
//
//1 是丑数。
//n 不超过1690。

func main() {
	fmt.Println(nthUglyNumber(10))
}

// 动态规划
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	a, b, c := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
