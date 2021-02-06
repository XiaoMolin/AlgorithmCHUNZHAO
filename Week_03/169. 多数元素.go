package main

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
//示例 1：
//
//输入：[3,2,3]
//输出：3
//示例 2：
//
//输入：[2,2,1,1,1,2,2]
//输出：2
//
//
//进阶：
//
//尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

func main() {

}

// 使用map
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for v, time := range m {
		if time > len(nums)/2 {
			return v
		}
	}
	return -1
}

// 使用计数器
func majorityElement1(nums []int) int {
	res, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			count++
		} else {
			count--
		}
	}
	return res
}

// 使用快速排序
func majorityElement2(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)/2]
}
func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}
