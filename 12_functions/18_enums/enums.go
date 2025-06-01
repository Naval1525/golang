package main

import "fmt"

//enumerated type
type orderStatus int
const (
	Received orderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status orderStatus){
	fmt.Println("Chaning order status to",status)
}
func main(){
	changeOrderStatus(Received)
}