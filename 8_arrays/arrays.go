package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println(len(nums))
	for i := 0; i < len(nums); i++ {
		nums[i] = i
		fmt.Println(nums[i])
	}
	nums_1 := [4]int{1, 2, 3, 4}
	fmt.Println(nums_1)

	//2d
	nums_2d := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(nums_2d)

	// fixed size, memory optimization , constant time access
}
