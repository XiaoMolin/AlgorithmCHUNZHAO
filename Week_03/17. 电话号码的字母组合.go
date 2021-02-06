package main

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//示例 1：
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//示例 2：
//
//输入：digits = ""
//输出：[]
//示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]

func main() {

}

// 使用一个表来记录字母和数字
// 对每一次字母求覆盖长度
// 时间复杂度O(n*m)
func letterCombinations(digits string) []string {
	str := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
	if len(digits) == 0 {
		return []string{}
	}
	length := 1
	for _, v := range digits {
		if v == '7' || v == '9' {
			length *= 4
			continue
		}
		length *= 3
	}
	res := make([]string, length)
	l := length
	for _, v := range digits {
		index := v - '0' - 2
		l /= len(str[index])
		i := 0
		k := 0
		for j := 0; j < length; j++ {
			k++
			res[j] = res[j] + str[index][i]
			if k%l == 0 {
				i++
				if i == len(str[index]) {
					i = 0
				}
			}
		}
	}
	return res
}
