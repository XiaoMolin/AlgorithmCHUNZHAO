package main

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	//使用双指针
	//将所有非零元素往前移，记录0的个数
	//对尾部进行归零处理
	j := 0
	time := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		} else {
			time++
		}
	}
	for ; time > 0; time-- {
		nums[len(nums)-time] = 0
	}
}
