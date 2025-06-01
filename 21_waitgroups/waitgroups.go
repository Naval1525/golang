package main

import (
	"fmt"
	"sync"
)

//wait groups
func task (id int,w *sync.WaitGroup){
	fmt.Println("doing task",id)
	defer w.Done() // signal that the task is done
}
func main(){
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i,&wg) // start a goroutine for each task

	}
	wg.Wait()
	fmt.Println("All tasks completed")


}