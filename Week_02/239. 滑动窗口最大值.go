package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

//双端队列
//下标入队
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	q := []int{}
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if q[0] == i-k {
			q = q[1:]
		}
		if i-k+1 >= 0 {
			res[i-k+1] = nums[q[0]]
		}
	}
	return res
}

//type Heap []int
//
//func (h Heap)Len()int {return len(h)}
//func (h Heap)Less(i,j int)bool {return h[i]>h[j]}
//func (h Heap)Swap(i,j int) {h[i],h[j]=h[j],h[i]}
//func (h *Heap)Push(x interface{}){
//	*h=append(*h,x.(int))
//}
//func (h *Heap)Pop()interface{}{
//	old :=*h
//	n:=len(old)
//	x :=old[n-1]
//	*h=old[:n-1]
//	return x
//}

// 使用大根堆
//func maxSlidingWindow2(nums []int, k int) []int {
//	h:=make(Heap,0)
//	res:=make([]int,len(nums)-k+1)
//	for i:=0;i<len(nums);i++{
//		start :=i-k
//		if start>=0{
//			heap.Remove(&h,nums[start])
//		}
//		heap.Push(&h,nums[i])
//		if h.Len()==k{
//			n:=heap.Pop(&h).(int)
//			res[i-k+1]=n
//			heap.Push(&h,n)
//		}
//
//	}
//	return res
//}
