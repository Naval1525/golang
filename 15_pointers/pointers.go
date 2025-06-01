package main

import "fmt"

//by reference
func c(num *int){
	*num = 5
	fmt.Println("in c",num)
}
func main(){
	num:=1
	c(&num)
	fmt.Println("in main",num) // in main 1

}