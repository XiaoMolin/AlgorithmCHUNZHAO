package main

//给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
//每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
//
//示例 1：
//
//输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//输出：11
//解释：如下面简图所示：
//2
//3 4
//6 5 7
//4 1 8 3
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//示例 2：
//
//输入：triangle = [[-10]]
//输出：-10
//
//
//提示：
//
//1 <= triangle.length <= 200
//triangle[0].length == 1
//triangle[i].length == triangle[i - 1].length + 1
//-104 <= triangle[i][j] <= 104

func main() {

}

// 自底向上
/*
2
3 4
5 6 7
2 3 1 2
*/
// dp[i][j]=min(dp[i+1][j],dp[i+1][j+1])+triangle[i][j]
func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
