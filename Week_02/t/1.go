package main

import "fmt"

func digui(a []int, i int) []int {
	copy(a[i:], a[i+1:])
	return a
}

func main() {
	a := []int{4, 6, 7, 8, 9}
	fmt.Println(digui(a, 2))
}
