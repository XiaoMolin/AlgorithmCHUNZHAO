package main

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//
//
//示例 1：
//
//输入：grid = [
//["1","1","1","1","0"],
//["1","1","0","1","0"],
//["1","1","0","0","0"],
//["0","0","0","0","0"]
//]
//输出：1
//示例 2：
//
//输入：grid = [
//["1","1","0","0","0"],
//["1","1","0","0","0"],
//["0","0","1","0","0"],
//["0","0","0","1","1"]
//]
//输出：3
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 300
//grid[i][j] 的值为 '0' 或 '1'

import (
	"fmt"
)

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1'}, {'0', '0'}}))
}

//  dfs 遇到1改为0 直到遍历结束
func numIslands(grid [][]byte) int {
	len1 := len(grid)
	if len1 == 0 {
		return 0
	}
	len2 := len(grid[0])
	res := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid)
				res++
			}
		}
	}
	fmt.Println(grid)
	return res
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)

}
