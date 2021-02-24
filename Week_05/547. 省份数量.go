package main

//有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
//
//省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
//
//给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
//
//返回矩阵中 省份 的数量。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-provinces
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// 使用一个数组模拟map
// 使用深度优先(非递归版本)
func findCircleNum(M [][]int) int {
	count := 0
	queue := []int{}
	visit := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		if visit[i] == 0 {
			queue = append(queue, i)
			for len(queue) != 0 {
				city := queue[0]
				queue = queue[1:]
				for k := 0; k < len(M[city]); k++ {
					if M[city][k] == 1 && visit[k] == 0 {
						visit[k] = 1
						queue = append(queue, k)
					}
				}
			}
			count++
		}
	}
	return count
}

//并查集
func findCircleNum1(isConnected [][]int) int {
	n := len(isConnected)
	if n == 1 {
		return 1
	}

	uf := generate(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

type unionFind struct {
	count  int
	father []int
}

func generate(n int) (uf *unionFind) {
	uf = &unionFind{}
	uf.count = n
	uf.father = make([]int, n)
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}
	return uf
}

func (uf *unionFind) union(p, q int) {
	if p == q {
		return
	}
	fq := uf.find(q)
	fp := uf.find(p)
	if fq == fp {
		return
	}
	uf.father[fp] = fq
	uf.count--
}

func (uf *unionFind) find(x int) int {
	r := x
	// 找到x的根
	for uf.father[r] != r {
		r = uf.father[r]
	}
	//将路径上的节点的值设置为根的值
	for uf.father[x] != x {
		t := uf.father[x]
		uf.father[t] = r
		x = t
	}
	return x
}
