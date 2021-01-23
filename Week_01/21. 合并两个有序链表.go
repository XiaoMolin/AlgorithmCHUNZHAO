package main

func main() {
	l1 := &ListNode{}
	l2 := &ListNode{}
	mergeTwoLists(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//类似归并排序 的第二步骤
	list := &ListNode{}
	cur := list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			c := &ListNode{
				Val: l1.Val,
			}
			cur.Next = c
			cur = cur.Next
			l1 = l1.Next
		} else {
			c := &ListNode{
				Val: l2.Val,
			}
			cur.Next = c
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return list.Next
}

func mergeTwoList2(l1 *ListNode, l2 *ListNode) *ListNode {
	//类似归并排序 的第二步骤
	//递归
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoList2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList2(l1, l2.Next)
		return l2
	}
}
