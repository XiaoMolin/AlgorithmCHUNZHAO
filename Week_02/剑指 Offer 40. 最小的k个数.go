package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(getLeastNumbers([]int{3, 4, 5, 1, 2, 3, 5, 7}, 3))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getLeastNumbers(arr []int, k int) []int {
	var res []int
	if len(arr) == 0 || k == 0 {
		return res
	}

	h := IntHeap(arr)
	h2 := &IntHeap{}
	h2 = &h
	heap.Init(h2)
	for k > 0 {
		res = append(res, heap.Pop(h2).(int))
		k--
	}

	return res

}
