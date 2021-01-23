package main

func main() {
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

func removeDuplicates(nums []int) int {
	//因为已经排好序，那么使用直接遍历即可
	if len(nums) == 0 {
		return 0
	}
	temp := nums[0]
	for i := 1; i < len(nums); {
		if nums[i] == temp {
			//重复
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			temp = nums[i]
			i++
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	//双指针
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
