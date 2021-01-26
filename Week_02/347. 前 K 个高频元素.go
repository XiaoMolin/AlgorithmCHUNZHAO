package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{3, 4, 6, 1, 23, 2, 5, 1}, 3))
}

//思路1
//将出现的所有数字存入maps里，个数为key
//示例 1:nums = [1,1,1,2,2,3]
//maps[1]=3 maps[2]=2 maps[3]=1
//然后使用堆排序维持堆的长度为topk(tokp为k或maps的长度取最小)
//然后将堆的num取出
//每次都将新的元素与堆顶元素（堆中频率最小的元素）进行比较
//如果新的元素的频率比堆顶端的元素大，则弹出堆顶端的元素，
//将新的元素添加进堆中
//最终，堆中的topk个元素即为前topk个高频元素
//
func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}
	maps := make(map[int]int)
	for _, v := range nums {
		maps[v]++
	}
	h := &NodeHeap{}
	topK := min(k, len(maps))
	size := 0
	for k, v := range maps {
		if size < topK {
			heap.Push(h, &Node{
				val: k,
				num: v,
			})
			size++
		} else {
			if v > (*h)[0].num {
				heap.Pop(h)
				heap.Push(h, &Node{
					val: k,
					num: v,
				})
			}
		}
	}
	res := make([]int, 0, topK)
	for i := 0; i < topK; i++ {
		res = append(res, heap.Pop(h).(*Node).val)
	}
	return res
}

type Node struct {
	val int
	num int
}
type NodeHeap []*Node

func (this NodeHeap) Len() int { return len(this) }

// 小根堆
func (this NodeHeap) Less(i, j int) bool { return this[i].num < this[j].num }

func (this NodeHeap) Swap(i, j int) { this[i], this[j] = this[j], this[i] }

func (this *NodeHeap) Push(x interface{}) {
	*this = append(*this, x.(*Node))
}

func (this *NodeHeap) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[:n-1]
	return x
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
