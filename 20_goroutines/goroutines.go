package main

import (
	"fmt"
	"time"
)

// func task (id int){
// 	fmt.Println("doing task",id)
// }
func main(){
	for i := 0; i < 10; i++ {
		// go task(i) // start a goroutine for each task
		go func(i int){
			fmt.Println(i)
		}(i)
	}
	fmt.Println("All tasks started")
	time.Sleep(time.Second*2)

}