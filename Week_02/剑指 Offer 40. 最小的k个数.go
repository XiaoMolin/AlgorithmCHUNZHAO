package main

//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
//示例 1：
//
//输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//示例 2：
//
//输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//限制：
//
//0 <= k <= arr.length <= 10000
//0 <= arr[i] <= 10000

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
