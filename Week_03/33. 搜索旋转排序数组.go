package main

//升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
//
//请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//
//
//示例 1：
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//示例 3：
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//提示：
//
//1 <= nums.length <= 5000
//-10^4 <= nums[i] <= 10^4
//nums 中的每个值都 独一无二
//nums 肯定会在某个点上旋转
//-10^4 <= target <= 10^4

func main() {

}

// 二分法 分两种情况
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// 左边单调递增
		if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右边单调递增
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}
