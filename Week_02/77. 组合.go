package main

import "fmt"

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

func main() {
	fmt.Println(combine1(2, 3))
}

//使用回溯
func combine1(n int, k int) [][]int {
	res := [][]int{}

	var dfs func(n, k, start int, path []int)
	dfs = func(n, k, start int, path []int) {
		if len(path) == k {
			res = append(res, path)
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			ll := make([]int, len(path))
			copy(ll, path)
			dfs(n, k, i+1, ll)
			path = path[:len(path)-1]
		}
	}
	return res
}

//效率高的回溯
func combine(n int, k int) [][]int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	res = [][]int{}
	recursion(nums, k, []int{})

	return res
}

var res [][]int

// 使用数组来截取
func recursion(nums []int, k int, path []int) {
	if len(path) == k {
		b := make([]int, k)
		copy(b, path)
		res = append(res, b)
		return
	}

	// 剪枝
	if k > len(path)+len(nums) {
		return
	}

	for i, top := range nums {
		numsNew := nums[i+1:]
		path = append(path, top)
		//fmt.Println(path,nums)
		recursion(numsNew, k, path)
		path = path[:len(path)-1]
	}
}
