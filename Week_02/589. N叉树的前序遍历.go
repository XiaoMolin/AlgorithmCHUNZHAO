package main

func main() {
	preorder(&Node{})
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := []int{}
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root != nil {
			res = append(res, root.Val)
			for _, node := range root.Children {
				dfs(node)
			}
		}
	}
	dfs(root)
	return res
}
