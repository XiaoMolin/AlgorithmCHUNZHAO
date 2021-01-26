package main

import (
	"fmt"
)

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
