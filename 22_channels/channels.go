package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func processNum(numChan chan int) {
	for {
		num := <-numChan
		fmt.Println("Processing number:", num)
		time.Sleep(time.Second)
	}
}

func main() {
	numChan := make(chan int)

	go processNum(numChan)

	for {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		numChan <- int(nBig.Int64())
		time.Sleep(500 * time.Millisecond) // throttle generation
	}
}
