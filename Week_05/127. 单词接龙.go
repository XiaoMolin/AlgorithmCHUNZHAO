package main

//字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
//
//序列中第一个单词是 beginWord 。
//序列中最后一个单词是 endWord 。
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典 wordList 中的单词。
//给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。

func main() {

}

// 双向bfs
// 每一次使用短的队列
func ladderLength(beginWord string, endWord string, wordList []string) int {
	exists := false
	for _, w := range wordList {
		if w == endWord {
			exists = true
		}
	}
	if !exists {
		return 0
	}
	var res int

	queue := []string{beginWord}
	visited := map[string]struct{}{beginWord: struct{}{}}

	queue1 := []string{endWord}
	visited1 := map[string]struct{}{endWord: struct{}{}}
	for len(queue) > 0 && len(queue1) > 0 {
		if len(queue) > len(queue1) {
			queue, queue1 = queue1, queue
			visited, visited1 = visited1, visited
		}
		n := len(queue)
		res++
		for i := 0; i < n; i++ {
			node := queue[i]
			for _, w := range wordList {
				//访问过
				if _, ok := visited[w]; ok {
					continue
				}
				//不能转换
				if !connect(node, w) {
					continue
				}
				//双向遇到
				if _, ok := visited1[w]; ok {
					return res + 1
				}
				queue = append(queue, w)
				visited[w] = struct{}{}
			}
		}
		queue = queue[n:]
	}
	return 0
}

func connect(w1, w2 string) bool {
	var cnt int
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			cnt++
		}
	}
	return cnt == 1
}
