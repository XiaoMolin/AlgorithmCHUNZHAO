package main

func main() {
	postorder(&Node{})
}

type Node struct {
	Val      int
	Children []*Node
}

//递归
func postorder(root *Node) []int {
	var res []int
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root != nil {
			for _, node := range root.Children {
				dfs(node)
			}
			res = append(res, root.Val)
		}
	}
	dfs(root)
	return res
}

//迭代
//广度优先遍历
//一层一层遍历，先入栈---后出栈--然后进入队列
func postorder1(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*Node{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val}, res...)
		for _, n := range node.Children {
			stack = append(stack, n)
		}
	}
	return res
}
