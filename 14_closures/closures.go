package main

func counter() func() int{
	var count int = 0
	return func() int {
		count++
		return count
	}
}
func main(){
	c1 := counter()
	println("Counter 1:", c1()) // Counter 1: 1
	println("Counter 1:", c1()) // Counter 1: 2
	println("Counter 1:", c1()) // Counter 1: 3

	c2 := counter()
	println("Counter 2:", c2()) // Counter 2: 1
	println("Counter 2:", c2()) // Counter 2: 2
	println("Counter 2:", c2()) // Counter 2: 3

	println("Counter 1:", c1()) // Counter 1: 4
	println("Counter 2:", c2()) // Counter 2: 4

}