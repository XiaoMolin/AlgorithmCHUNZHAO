package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]

func main() {

}

// 递归
//一次一半阔号
//使用回溯
//如果左括号数量不大于 nn，我们可以放一个左括号。
//如果右括号数量小于左括号的数量，我们可以放一个右括号
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	outs := make([]string, 0)
	dfs(0, 0, n, "", &outs)
	return outs
}

func dfs(left, right, n int, str string, outs *[]string) {
	if left == right && left == n {
		*outs = append(*outs, str)
		return
	}
	if left < n {
		dfs(left+1, right, n, str+"(", outs)
	}
	if right < left {
		dfs(left, right+1, n, str+")", outs)
	}
}
