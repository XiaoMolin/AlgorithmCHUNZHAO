package main

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[string]int)
	for i := 0; i < len(s); i++ {
		maps[string(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		maps[string(t[i])]--
		if maps[string(t[i])] < 0 {
			return false
		}
	}
	return true
}
