package main

import "fmt"

func main() {
	nums := []int{3, 5, 1, 2, 7, 8}
	fmt.Println(merge(nums))

}

func mergeSort(nums1, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	res := []int{}
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	res = append(res, nums1[i:]...)
	res = append(res, nums2[j:]...)
	return res
}

func merge(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	med := len(nums) / 2
	nums1 := merge(nums[:med])
	nums2 := merge(nums[med:])
	return mergeSort(nums1, nums2)
}
