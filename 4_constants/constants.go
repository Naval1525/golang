package main

import "fmt"

const (
	port = 5000
	host = "localhost"
)

func main() {
	// constant cannot be changed
	const name = "Naval"
	const age = 30
	const isAdult = true
	fmt.Println(port, host)
}
