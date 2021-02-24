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

func main() {

}

//并查集
//并查集
func numIslands(grid [][]byte) int {
	t := 0
	n := len(grid)*len(grid[0]) + 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				t++
			}
		}
	}
	var position func(i, j int) int
	position = func(i, j int) int {
		return i*len(grid[0]) + j
	}
	u := NewUnionFind(n, t)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i > 0 && grid[i][j] == '1' && grid[i-1][j] == '1' {
				u.Union(position(i, j), position(i-1, j))
			}
			if j > 0 && grid[i][j] == '1' && grid[i][j-1] == '1' {
				u.Union(position(i, j), position(i, j-1))
			}
		}
	}
	return u.count
}

type UninoFind struct {
	parent []int
	count  int
}

func NewUnionFind(n, t int) *UninoFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UninoFind{
		parent: parent,
		count:  t,
	}
}

func (u *UninoFind) Union(i, j int) {
	pi, pj := u.Find(i), u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UninoFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
