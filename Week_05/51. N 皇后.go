package main

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

func main() {

}

func solveNQueens(n int) [][]string {
	res = [][]string{}

	board := make([][]byte, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]byte, n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = '.'
		}
	}
	dfs(board, n, 0)
	return res
}

var res [][]string

func dfs(board [][]byte, n int, deep int) {
	// fmt.Println(n,deep,board,cur)
	if n == deep {
		// fmt.Println(n)
		boardByteToString(board)
		return
	}
	for i := 0; i < n; i++ {
		if isValid(board, deep, i) {
			board[deep][i] = 'Q'
			dfs(board, n, deep+1)
			board[deep][i] = '.'
		}
	}
}

func isValid(board [][]byte, row, col int) bool {
	//检查行是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	//检查列是否有冲突
	for i := 0; i < len(board); i++ {
		if board[row][i] == 'Q' {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
func boardByteToString(board [][]byte) {
	rr := []string{}
	for i := 0; i < len(board); i++ {
		r := ""
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'Q' {
				r += "Q"
			} else {
				r += "."
			}
		}
		rr = append(rr, r)
	}
	res = append(res, rr)
}
