package main

//给定一个二叉树，返回它的 后序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [3,2,1]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

func main() {
	postorderTraversal(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 普通迭代
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	pre := &TreeNode{}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) || (pre == cur.Left || pre == cur.Right) {
			res = append(res, cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}

	}
	return res
}

// 颜色标记法
func postorderTraversal2(root *TreeNode) []int {

	white, gray := 0, 1
	res := []int{}
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
			//左右跟
			stack = append(stack, s{gray, node})
			stack = append(stack, s{white, node.Right})
			stack = append(stack, s{white, node.Left})
		} else {
			res = append(res, node.Val)
		}
	}
	return res
}
