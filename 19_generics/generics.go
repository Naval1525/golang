package main

import "fmt"

// //generic
// func printSlice[T int | string](items []T){
// 	for i, item := range items {
// 		println("item", i, ":", item)
// 	}

// }
//LIFO
type Stack[T any] struct {
	elements []T
}
func main() {
	// slice := []int{1, 2, 3, 4, 5}
	// printSlice(slice)
	// printSlice(slice)
	myStack := Stack[string]{
		elements: []string{"Hello","Buddy"},
	}
	fmt.Println(myStack)

}
