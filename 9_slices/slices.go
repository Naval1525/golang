package main

import (
	"fmt"
	"slices"
)

// slices are dynamic array
// most used construct in go
func main() {
	//uninialized slice is nil (same to null)
	var nums []int
	fmt.Println(nums)
	//for not to be null
	var num = make([]int, 2, 2)
	fmt.Println(cap(num))
	num = append(num, 1)
	num = append(num, 2)
	fmt.Println(cap(num))
	//copy
	var num2 = make([]int, len(num))
	//copy
	copy(num2, num)
	fmt.Println(num)
	fmt.Println(num2)

	//slice operator -> for range
	var nums2 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums2[0:3])

	//slice
	var n1 = []int{1, 2, 3}
	var n2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(n1, n2))
	//2d slices
	var nums2d = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(nums2d[1][2]) // Accessing element at row 1, column 2

}
