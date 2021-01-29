package main

import "math"

//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明：叶子节点是指没有子节点的节点。
//
//
//
//示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：2
//示例 2：
//
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5

func main() {
	minDepth(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归求解
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	mind := math.MaxInt32
	if root.Left != nil {
		mind = min(minDepth(root.Left), mind)
	}
	if root.Right != nil {
		mind = min(minDepth(root.Right), mind)
	}
	return mind + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
