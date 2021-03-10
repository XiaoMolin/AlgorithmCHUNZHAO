package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//func main() {
//	nums:=[]int{3,5,1,2,7,8}
//	quickSort(nums,0,len(nums)-1)
//	fmt.Println(nums)
//
//}
// 1单函数
func quickSort1(nums []int, left, right int) {
	if left < right {
		key, i, j := nums[left], left, right
		for {
			for nums[i] < key {
				i++
			}
			for nums[j] > key {
				j--
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		quickSort(nums, left, i-1)
		quickSort(nums, j+1, right)
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
func quickSort_go(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSort_go(a, lo, p-1, childDone, depth)
		go quickSort_go(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
	done <- struct{}{}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3, testData4 := make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000)
	times := 100000000
	for i := 0; i < times; i++ {
		val := rand.Intn(20000000)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
		testData4 = append(testData3, val)
	}
	start := time.Now()
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("single goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}
	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(testData2, 0, len(testData2)-1, done, 5)
	<-done
	fmt.Println("multiple goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quickSort_go implementation")
	}
	start = time.Now()
	sort.Ints(testData3)
	fmt.Println("std lib: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong std lib implementation")
	}
	start = time.Now()
	quickSort1(testData4, 0, len(testData4)-1)
	fmt.Println("single left partition: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong single left partition implementation")
	}
}
