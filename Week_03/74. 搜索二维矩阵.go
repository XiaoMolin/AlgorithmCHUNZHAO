package main

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//
//
//示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 100
//-104 <= matrix[i][j], target <= 104

func main() {

}

// 二分查找
// 下标使用 i*len(matrix)+j
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if len(matrix) == 0 || len(matrix[0]) == 0 || target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	}
	l := 0
	r := ((len(matrix)) * (len(matrix[0]))) - 1
	for l <= r {
		mid := l + (r-l)>>1
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l > m*n || matrix[l/n][l%n] != target {
		return false
	}
	return true
}
