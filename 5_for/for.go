package main

import "fmt"

// for -> only for-loop in go no while and all
func main() {
	//while loop
	i := 0
	for i <= 5 {
		println(i)
		i++
	}
	//infinite loop
	for {
		fmt.Println("Infinite Loop")
		break
	}
	//classic for-loop
	for j := 0; j <= 3; j++ {
		if i == 2 {
			continue
		}
		fmt.Println(j)
	}
	//range loop
	for i := range 3 {
		fmt.Println(i)
	}

}
