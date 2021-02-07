package main

import (
	"math"
)

//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换后得到的单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回一个空列表。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出:
//[
//["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

func main() {

}

// 使用广度优先
//
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := [][]string{}
	edFlag := false
	for _, word := range wordList {
		if word == endWord {
			edFlag = true
			break
		}
	}
	if !edFlag {
		return res
	}
	bfs(beginWord, endWord, wordList, &res)
	return res
}

func bfs(beginWord string, endWord string, wordList []string, res *[][]string) {
	queue := [][]string{}
	path := []string{beginWord}
	queue = append(queue, path)
	isFound := false
	dict := make(map[string]bool)
	for _, word := range wordList {
		dict[word] = true
	}
	visited := make(map[string]bool)
	visited[beginWord] = true
	for len(queue) > 0 {
		size := len(queue)
		subVisited := make(map[string]bool)
		for j := 0; j < size; j++ {
			p := queue[0]
			queue = queue[1:]
			//得到当前路径的末尾单词
			temp := p[len(p)-1]
			// 一次性得到所有的下一个的节点
			neighborus := getNeighbors(temp, dict)
			for _, neighbor := range neighborus {
				//只考虑之前没有出现过的单词
				if !visited[neighbor] {
					if neighbor == endWord {
						isFound = true
						p = append(p, neighbor)
						*res = append(*res, append([]string(nil), p...))
						p = p[:len(p)-1]
					}
					//加入当前单词
					p = append(p, neighbor)
					queue = append(queue, append([]string(nil), p...))
					p = p[:len(p)-1]
					subVisited[neighbor] = true
				}
			}
		}
		for s, _ := range subVisited {
			visited[s] = true
		}
		if isFound {
			break
		}
	}
}

func getNeighbors(node string, dict map[string]bool) []string {
	res := []string{}
	chs := []byte(node)
	for ch := 'a'; ch <= 'z'; ch++ {
		for i := 0; i < len(chs); i++ {
			if chs[i] == byte(ch) {
				continue
			}
			oldCh := chs[i]
			chs[i] = byte(ch)
			if dict[string(chs)] {
				res = append(res, string(chs))
			}
			chs[i] = oldCh
		}
	}
	return res
}

// 先转换成图
func findLadders2(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}
	//添加begin到单词表
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, len(wordList))
	// 建立邻接矩阵
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isMissOne(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}
	res := [][]string{}
	cost := make([]int, n)
	queue := [][]int{[]int{ids[beginWord]}}
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

func isMissOne(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i+1:] == b[i+1:]
		}
	}
	return false
}

// 超时 bfs+dfs
//var length int
//var res [][]string
//var map2 map[string]int
//func findLadders(beginWord string, endWord string, wordList []string) [][]string {
//	wordMap := map[string]bool{}
//	res=[][]string{}
//	for _, w := range wordList {
//		wordMap[w] = true
//	}
//	if _,ok:=wordMap[endWord];!ok{
//		return [][]string{}
//	}
//	length=1
//	bfs(beginWord,endWord,wordList,wordMap)
//	if length==1{
//		return [][]string{}
//	}
//	for _, w := range wordList {
//		wordMap[w] = true
//	}
//	map2=map[string]int{}
//	for _, w := range wordList {
//		map2[w] = math.MaxInt32
//	}
//	dfs(beginWord,endWord,wordList,wordMap,[]string{beginWord})
//	return res
//}
//
//func bfs(beginWord,endWord string,wordList []string,wordMap map[string]bool)int{
//	queue := []string{beginWord}
//	for len(queue) != 0 {
//		s := len(queue)
//		for i := 0; i < s; i++ {
//			word := queue[0]
//			queue = queue[1:]
//			if word == endWord {
//				return length
//			}
//			for c := 0; c < len(word); c++ {
//				for j := 'a'; j <= 'z'; j++ {
//					newWord := word[:c] + string(j) + word[c+1:]
//					if wordMap[newWord] == true {
//						queue = append(queue, newWord)
//						wordMap[newWord] = false
//					}
//				}
//			}
//		}
//		length++
//	}
//	return 0
//}
//
//func dfs(beginWord,endWord string, wordList []string,wordMap map[string]bool,path []string){
//	if len(path)==length&&path[len(path)-1]==endWord{
//		tmp:=make([]string,len(path))
//		copy(tmp,path)
//		res=append(res,tmp)
//	}else if len(path)>=length&&path[len(path)-1]!=endWord{
//		return
//	}
//	for i:=0;i<len(wordList);i++{
//		if isMissOne(wordList[i],beginWord)&&wordMap[wordList[i]]&&len(path)<=map2[wordList[i]]{
//			wordMap[wordList[i]]=false
//			map2[wordList[i]]=len(path)
//			path=append(path,wordList[i])
//			dfs(wordList[i],endWord,wordList,wordMap,path)
//			path=path[:len(path)-1]
//			wordMap[wordList[i]]=true
//		}
//		fmt.Println(path)
//
//	}
//
//}
