package main

import (
	"fmt"
	"sync"
)

type post struct {
	//race condition se bachne kailiye mutex use karte hai
	views int
	//mutex ko struct mai he rkha jata hai
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()

	}()

	//sirf lock whi lagao jha jarurat hooo
	p.mu.Lock()
	p.views += 1

}
func main() {
	myPost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)

}
