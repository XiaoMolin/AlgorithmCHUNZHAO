package main

import "sort"

func main() {
	isAnagram("asd", "asds")
}

//排序后对比
func isAnagram(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	for i, v := range s1 {
		if v != t1[i] {
			return false
		}
	}
	return true
}

//使用哈希表
func isAnagram2(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	maps := make(map[byte]int)
	for i := 0; i < l1; i++ {
		maps[s[i]]++
	}
	for i := 0; i < l2; i++ {
		maps[t[i]]--
		if maps[t[i]] < 0 {
			return false
		}
	}
	return true
}
