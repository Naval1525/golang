package main

import (
	"fmt"
	"time"
)

//order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
}
func main(){
	order := order{
		id: "12345",
		amount: 100.50,
		status: "pending",
		createdAt: time.Now(),
	}
	fmt.Println("Order ID:", order)

}

