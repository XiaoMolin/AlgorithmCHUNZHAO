package main

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶

func main() {

}

// dfs超时
func climbStairs(n int) int {
	res := 0
	var dfs func(i int)
	dfs = func(i int) {
		if i > n {
			return
		}
		if i == n {
			res++
		}
		dfs(i + 1)
		dfs(i + 2)
	}
	dfs(0)
	return res
}

func climbStairs(n int) int {
	prev := 1
	cur := 1
	for i := 2; i < n+1; i++ {
		temp := cur
		cur = prev + cur
		prev = temp
	}
	return cur
}
