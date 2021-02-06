package main

import "math"

//您需要在二叉树的每一行中找到最大的值。
//
//示例：
//
//输入:
//
//1
/// \
//3   2
/// \   \
//5   3   9
//
//输出: [1, 3, 9]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 使用迭代
// 时间复杂度O(n) n为节点数
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		max := queue[0].Val
		for i := 0; i < length; i++ {
			temp := queue[i]
			if temp.Val > max {
				max = temp.Val
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		queue = queue[length:]
		res = append(res, max)
	}

	return res
}

// 递归
func largestValues2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, math.MinInt32)
	}
	(*res)[level] = int(math.Max(float64((*res)[level]), float64(root.Val)))
	if root.Left != nil {
		dfs(root.Left, res, level+1)
	}
	if root.Right != nil {
		dfs(root.Right, res, level+1)
	}
}
