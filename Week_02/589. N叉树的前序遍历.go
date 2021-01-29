package main

//给定一个 N 叉树，返回其节点值的前序遍历。
//
//例如，给定一个 3叉树 :
//
//
//
//
//
//
//
//返回其前序遍历: [1,3,5,6,2,4]。
//
//
//
//说明: 递归法很简单，你可以使用迭代法完成此题吗?

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
