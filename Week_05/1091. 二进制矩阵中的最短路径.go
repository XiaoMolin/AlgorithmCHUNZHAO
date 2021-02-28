package main

//给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
//
//二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：
//
//路径途经的所有单元格都的值都是 0 。
//路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
//畅通路径的长度 是该路径途经的单元格总数。

func main() {

}

// 使用A* Search
// 参考题解实现
func shortestPathBinaryMatrix(grid [][]int) int {
	h := len(grid)
	if grid[0][0] == 1 || grid[h-1][h-1] == 1 {
		return -1
	}
	if h == 1 {
		return 1
	}
	// 优先函数
	getPriority := func(p [2]int, step int) int {
		return max(h-1-p[0], h-1-p[1]) + step*4
	}
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, h)
	}
	pq := &heap{[]kv{
		{[2]int{0, 0}, 1, getPriority([2]int{0, 0}, 1)},
	}}
	used[0][0] = true
	for len(pq.l) != 0 {
		node := pq.Pop()
		for _, p := range get8Conn(node.p) {
			// 找到最右下
			if p[0] == h-1 && p[1] == h-1 {
				return node.step + 1
			}
			if p[0] < 0 || p[0] > h-1 || p[1] < 0 || p[1] > h-1 || used[p[0]][p[1]] {
				continue
			}
			used[p[0]][p[1]] = true
			if grid[p[0]][p[1]] == 1 {
				continue
			}
			pq.Push(kv{
				p, node.step + 1, getPriority(p, node.step+1),
			})
		}
	}
	return -1
}

func get8Conn(p [2]int) [8][2]int {
	return [8][2]int{
		{p[0] + 1, p[1]}, {p[0] - 1, p[1]}, {p[0], p[1] + 1}, {p[0], p[1] - 1},
		{p[0] + 1, p[1] + 1}, {p[0] + 1, p[1] - 1}, {p[0] - 1, p[1] + 1}, {p[0] - 1, p[1] - 1},
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 以下是堆的实现，可自行实现
type kv struct {
	p     [2]int
	step  int
	value int
}

type heap struct {
	l []kv
}

func (h *heap) Push(i kv) {
	h.l = append(h.l, i)
	h.up(len(h.l) - 1)
}

func (h *heap) up(start int) {
	tmp := h.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if tmp.value >= h.l[par].value {
			break
		}
		h.l[cur] = h.l[par]
		cur = par
		par = (par - 1) / 2
	}
	h.l[cur] = tmp
}

func (h *heap) Top() kv {
	if len(h.l) == 0 {
		return kv{}
	}
	return h.l[0]
}

func (h *heap) Pop() kv {
	if len(h.l) == 0 {
		return kv{}
	}
	res := h.l[0]
	h.l[0] = h.l[len(h.l)-1]
	h.l = h.l[:len(h.l)-1]
	if len(h.l) != 0 {
		h.down(0)
	}
	return res
}

func (h *heap) down(start int) {
	tmp := h.l[start]
	cur := start
	child := cur*2 + 1
	for child < len(h.l)-1 {
		if h.l[child].value > h.l[child+1].value {
			child++
		}
		if tmp.value < h.l[child].value {
			break
		}
		h.l[cur] = h.l[child]
		cur = child
		child = child*2 + 1
	}
	h.l[cur] = tmp
}
