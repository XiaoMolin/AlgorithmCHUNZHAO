package main

//
//88. 合并两个有序数组
//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//示例 2：
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//
//
//提示：
//
//nums1.length == m + n
//nums2.length == n
//0 <= m, n <= 200
//1 <= m + n <= 200
//-109 <= nums1[i], nums2[i] <= 109

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

//思路1：双指针，先确定长度，然后从后往前填数
func merge(nums1 []int, m int, nums2 []int, n int) {
	right := m + n
	m = m - 1
	n = n - 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[right] = nums1[m]
			m--
			right--
		} else {
			nums1[right] = nums2[n]
			n--
			right--
		}
	}
	for n >= 0 {
		nums1[right] = nums2[n]
		n--
		right--
	}

}
