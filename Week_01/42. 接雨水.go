package main

import "fmt"

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
//示例 1：
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	left, right, leftMax, rightMax, res := 0, 0, 0, 0, 0
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//设置左边最高柱子
				leftMax = left
			} else {
				//右边有柱子挡水，所以，所有小于等于的都会积水
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				//设置右边最高柱子
				rightMax = height[right]
			} else {
				//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				res += rightMax - height[right]
			}
			right--
		}

	}
	return res
}
