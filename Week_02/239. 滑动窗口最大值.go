package main

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//
//
//
//示例 1：
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
//示例 2：
//
//输入：nums = [1], k = 1
//输出：[1]
//示例 3：
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//示例 4：
//
//输入：nums = [9,11], k = 2
//输出：[11]
//示例 5：
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//提示：
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104
//1 <= k <= nums.length

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
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
