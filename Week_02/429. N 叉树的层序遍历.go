package main

func main() {
	levelOrder(&Node{})
}

type Node struct {
	Val      int
	Children []*Node
}

// 使用栈

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	stack := []*Node{root}
	for len(stack) != 0 {
		n := len(stack)
		r := []int{}
		for i := 0; i < n; i++ {
			r = append(r, stack[i].Val)
			for _, n := range stack[i].Children {
				stack = append(stack, n)
			}
		}
		res = append(res, r)
		stack = stack[n:]
	}

	return res
}
