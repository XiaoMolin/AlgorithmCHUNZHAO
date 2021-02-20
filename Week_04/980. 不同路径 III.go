package main

//在二维网格 grid 上，有 4 种类型的方格：
//
//1 表示起始方格。且只有一个起始方格。
//2 表示结束方格，且只有一个结束方格。
//0 表示我们可以走过的空方格。
//-1 表示我们无法跨越的障碍。
//返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
//
//每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。
//
//
//
//示例 1：
//
//输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//输出：2
//解释：我们有以下两条路径：
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)

func main() {

}

// 只会dfs，想不出dp怎么写
var res int

func uniquePathsIII(grid [][]int) int {
	res = 0
	m := make([][]int, len(grid))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, len(grid[i]))
	}
	x, y := 0, 0
	z := 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				m[i][j] = 1
				x = i
				y = j
			}
			if grid[i][j] == -1 {
				z++
			}
		}
	}
	dfs(grid, x, y, m, z)
	return res
}

func dfs(grid [][]int, x, y int, m [][]int, num int) {
	if num == len(grid)*len(grid[0]) && grid[x][y] == 2 {
		res += 1
		return
	}
	xy := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(xy); i++ {
		x += xy[i][0]
		y += xy[i][1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != -1 && m[x][y] != 1 {
			m[x][y] = 1
			dfs(grid, x, y, m, num+1)
			m[x][y] = 0
		}
		x -= xy[i][0]
		y -= xy[i][1]
	}
}
