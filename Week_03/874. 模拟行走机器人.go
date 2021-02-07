package main

//机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：
//
//-2 ：向左转 90 度
//-1 ：向右转 90 度
//1 <= x <= 9 ：向前移动 x 个单位长度
//在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。
//
//机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
//
//返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）
//
//
//注意：
//
//北表示 +Y 方向。
//东表示 +X 方向。
//南表示 -Y 方向。
//西表示 -X 方向。
//
//
//示例 1：
//
//输入：commands = [4,-1,3], obstacles = []
//输出：25
//解释：
//机器人开始位于 (0, 0)：
//1. 向北移动 4 个单位，到达 (0, 4)
//2. 右转
//3. 向东移动 3 个单位，到达 (3, 4)
//距离原点最远的是 (3, 4) ，距离为 32 + 42 = 25
//示例 2：
//
//输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//输出：65
//解释：机器人开始位于 (0, 0)：
//1. 向北移动 4 个单位，到达 (0, 4)
//2. 右转
//3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
//4. 左转
//5. 向北走 4 个单位，到达 (1, 8)
//距离原点最远的是 (1, 8) ，距离为 12 + 82 = 65
//
//
//提示：
//
//1 <= commands.length <= 104
//commands[i] is one of the values in the list [-2,-1,1,2,3,4,5,6,7,8,9].
//0 <= obstacles.length <= 104
//-3 * 104 <= xi, yi <= 3 * 104
//答案保证小于 231

func main() {

}

// 按照题意
// 每走一次记录距离
func robotSim(commands []int, obstacles [][]int) int {
	obstacles_m := make(map[[2]int]bool)
	for _, v := range obstacles {
		if len(v) == 2 {
			obstacles_m[[2]int{v[0], v[1]}] = true
		}
	}
	//direct* = [4]int{north ,east, south, west}
	directForY := [4]int{1, 0, -1, 0}
	//y轴坐标四个方向的移动如何处理
	directForX := [4]int{0, 1, 0, -1}
	//x轴坐标四个方向的移动如何处理
	x := 0      //start x
	y := 0      //start y
	direct := 0 // 方向索引
	var distance int
	for len(commands) != 0 {
		command := commands[0]
		commands = commands[1:]
		if command == -1 {
			distance = max(distance, x*x+y*y)
			direct = (direct + 1) % 4
			continue
		}
		if command == -2 {
			distance = max(distance, x*x+y*y)
			direct = (direct + 3) % 4
			continue
		}

		for i := 0; i < command; i++ {
			tmpx := x + directForX[direct]*1
			tmpy := y + directForY[direct]*1
			if _, ok := obstacles_m[[2]int{tmpx, tmpy}]; ok {
				break
			}
			x = tmpx
			y = tmpy
		}
		//distance=max(distance,x*x+y*y)
	}
	return max(distance, x*x+y*y)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
