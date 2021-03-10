package main

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
func main() {

}

func reverseWords(s string) string {
	b := []byte(s)
	l := 0
	for i, v := range s {
		if v == ' ' || i == len(s)-1 {
			r := i - 1
			if i == len(s)-1 {
				r = i
			}
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}
			l = i + 1
		}
	}
	return string(b)
}
