package main

func main() {
	inorderTraversal(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归
// 使用一个栈按照先序遍历的顺序压栈
// 没有左节点就出栈，出栈的节点有右子树，右子树进栈
// 每一次出栈将结果放如数组
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

// 递归
// 按照访问左子树——根节点——右子树的方式遍历这棵树
func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, res []int) []int
	dfs = func(root *TreeNode, res []int) []int {
		if root.Left != nil {
			res = dfs(root.Left, res)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			res = dfs(root.Right, res)
		}
		return res
	}
	return dfs(root, res)
}

// 神奇的解法--- 颜色标记法 go语言的gc 采用的是三色标记法 有点类似
// 使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色
// 如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈
// 如果遇到的节点为灰色，则将节点的值输出
func inorderTraversal3(root *TreeNode) []int {
	white, gray := 0, 1
	var res []int
	type s struct {
		color int
		node  *TreeNode
	}
	stack := []s{{color: white, node: root}}
	for len(stack) != 0 {
		color, node := stack[len(stack)-1].color, stack[len(stack)-1].node
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		if color == white {
			stack = append(stack, s{white, node.Right})
			stack = append(stack, s{gray, node})
			stack = append(stack, s{white, node.Left})
		} else {
			res = append(res, node.Val)
		}
	}
	return res
}
