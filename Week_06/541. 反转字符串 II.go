package main

//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

func main() {

}

func reverseStr(s string, k int) string {
	if k == 1 || len(s) <= 1 {
		return s
	}

	in := []byte(s)
	cnt := len(in) / (2 * k)
	for i := 0; i < len(in)/(2*k); i++ {
		reverse(in[i*2*k : i*2*k+k])
	}

	if len(in)%(2*k) >= k {
		reverse(in[cnt*2*k : cnt*2*k+k])
	} else {
		reverse(in[cnt*2*k:])
	}

	return string(in)

}

func reverse(word []byte) {
	if len(word) <= 1 {
		return
	}

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return
}
