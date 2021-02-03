package main

import "fmt"

func digui(age int, num int) int {
	if num > 0 {
		age = digui(age+2, num-1)
	}
	return age
}

func main() {
	fmt.Println(digui(10, 1))
}
