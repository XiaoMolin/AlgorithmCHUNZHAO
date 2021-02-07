package main

//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
//说明：不要使用任何内置的库函数，如  sqrt。
//
//示例 1：
//
//输入：16
//输出：True
//示例 2：
//
//输入：14
//输出：False

func main() {

}

// 简单二分
func isPerfectSquare(num int) bool {
	l, r := 0, num/2
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return (l * l) == num
}
