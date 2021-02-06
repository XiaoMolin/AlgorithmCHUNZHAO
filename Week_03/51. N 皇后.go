package main

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//示例 2：
//
//输入：n = 1
//输出：[["Q"]]
//
//
//提示：
//
//1 <= n <= 9
//皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

func main() {

}

// 使用回溯法
// 每次放的时候监测是否冲突
// 时间复杂度O(n!)
var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	path := [][]byte{}
	backtrack(board, path)
	return res
}

func backtrack(board [][]bool, path [][]byte) {
	//结束条件
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, v := range path {
			t[k] = string(v)
		}
		res = append(res, t)
	}
	for i := 0; i < len(board); i++ {
		if !isValid(board, len(path), i) {
			continue
		}
		bs := printLine(len(board))
		bs[i] = 'Q'
		board[len(path)][i] = true
		path = append(path, bs)
		backtrack(board, path)
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

func isValid(board [][]bool, row, col int) bool {
	//检查行是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}
	//检查列是否有冲突
	for i := 0; i < len(board); i++ {
		if board[row][i] == true {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}
	return true
}

func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}
