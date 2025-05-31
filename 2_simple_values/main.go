package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2.5)
	fmt.Println("Hello World")
	fmt.Println(true)
	fmt.Println('A')                // rune, which is an alias for int32
	fmt.Println('B')                // rune, which is an alias for int32
	fmt.Println("Hello " + "World") // String concatenation
	fmt.Println("Hello " + "World" + "!")

}
