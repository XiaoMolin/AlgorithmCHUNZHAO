package main

//
//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//4
///   \
//2     7
/// \   / \
//1   3 6   9
//输出：
//
//4
///   \
//7     2
/// \   / \
//9   6 3   1
//备注:
//这个问题是受到 Max Howell 的 原问题 启发的 ：

func main() {
	invertTree(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
