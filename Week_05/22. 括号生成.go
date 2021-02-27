package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]

func main() {

}

//dfs
func generateParenthesis(n int) []string {
	res = []string{}

	generate(0, 0, n, "")
	return res
}

var res []string

func generate(left, right int, n int, s string) {
	if left == n && right == n {
		str := s
		res = append(res, str)
		return
	}
	if left < n {
		generate(left+1, right, n, s+"(")
	}
	if left < n && right < left {
		generate(left, right+1, n, s+")")
	}
}
