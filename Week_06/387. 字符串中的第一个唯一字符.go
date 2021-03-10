package main

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func main() {

}

func firstUniqChar(s string) int {
	var lf [26]int
	for i, ch := range s {
		lf[ch-'a'] = i
	}
	for i, ch := range s {
		if i == lf[ch-'a'] {
			return i
		} else {
			lf[ch-'a'] = -1
		}
	}
	return -1
}
