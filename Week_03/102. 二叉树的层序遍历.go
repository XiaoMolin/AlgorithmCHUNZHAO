package main

import "math"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层序遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 直接一个队列先进先出
// 时间复杂度O(n)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) != 0 {
		length := len(queue)
		r := []int{}
		for i := 0; i < length; i++ {
			temp := queue[i]
			r = append(r, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, r)
		queue = queue[length:]
	}
	return res
}

//递归
var res [][]int

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)

}
