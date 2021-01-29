package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]
//
//
//提示：
//
//1 <= n <= 8

func main() {

}

/*

一次添加一对括号 使用哈希表维护
思路：
1.
()
2.
()() (())
3.
()()() (())() ()()() ()(()) ()()()
()(()) (()()) ((())) (()()) (())()

然后使用map去除重复
可以优化成动态规划
*/
func generateParenthesis1(n int) []string {
	res := []string{}
	str := "()"
	dfs1(&res, str, n)
	m := make(map[string]int)
	for _, v := range res {
		m[v]++
	}
	result := []string{}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func dfs1(s *[]string, str string, n int) {
	if n == 1 {
		*s = append(*s, str)
		return
	}
	m := make(map[string]int)
	l := len(str)
	for i := 0; i < l; i++ {
		s1 := string(str[:i])
		s2 := string(str[i:])
		s3 := s1 + "()" + s2
		m[s3] = m[s3] + 1
		if m[s3] == 1 {
			dfs1(s, s3, n-1)
		}
	}
}

/*
一次一半阔号
使用回溯
如果左括号数量不大于 nn，我们可以放一个左括号。
如果右括号数量小于左括号的数量，我们可以放一个右括号
*/

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	outs := make([]string, 0)
	dfs(0, 0, n, "", &outs)
	return outs
}

func dfs(left, right int, n int, str string, outs *[]string) {
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
