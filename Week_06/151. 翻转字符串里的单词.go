package main

//给定一个字符串，逐个翻转字符串中的每个单词。
//
//说明：
//
//无空格字符构成一个 单词 。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

func main() {

}

func reverseWords(s string) string {
	maps := make(map[int]string)
	str := ""
	i := 1
	for _, v := range s {
		if v == ' ' && str != "" {
			maps[i] = str + " "
			str = ""
			i++
		} else if v != ' ' {
			str += string(v)
		}
	}
	if str != "" {
		maps[i] = str + " "
	}
	// fmt.Println(maps)
	res := ""
	for ; i >= 1; i-- {
		res += maps[i]

		// fmt.Println(res)
	}
	if len(maps) == 0 {
		return ""
	}
	return res[:len(res)-1]
}
