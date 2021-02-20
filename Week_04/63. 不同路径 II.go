package main

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//示例 1：
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//提示：
//
//m == obstacleGrid.length
//n == obstacleGrid[i].length
//1 <= m, n <= 100
//obstacleGrid[i][j] 为 0 或 1

func main() {

}

// 时间复杂度 O(nm)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	var dp [][]int
	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] != 1 { //如果起点不是障碍物，置为1，dp开始
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 { //当前点上有障碍物，位于障碍物位置的dp的值保持不变
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j] //当前点左侧有点
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1] //同上，上方的点
			}
		}
	}
	return dp[m-1][n-1]
}
