package main

func sum(nums ...int) int{
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main(){
	result := sum(1, 2, 3, 4, 5)
	println("The sum is:", result) // The sum is: 15
	result = sum(10, 20, 30)
	println("The sum is:", result) // The sum is: 60
	result = sum(100, 200)
	println("The sum is:", result) // The sum is: 300
	println("The sum is:", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // The sum is: 55

}