package main

import "fmt"

//interface
type paymenter interface{
	pay(amount float32)

}

//open close principle is not followed
type payment struct{
	gateway paymenter
}

func (p payment) makePayment (amount float32){
	p.gateway.pay(amount)

}


// type razorpay struct {}

type fakepay struct{}

func (f fakepay) pay(amount float32){
	fmt.Println("making payment using testing process", amount)
}


// func (r razorpay) pay(amount float32){
// 	fmt.Println("making payment using razorpay",amount)
// }


func main(){
	fake := fakepay{}
	newpayment := payment{
		gateway: fake,
	}
	newpayment.makePayment(100)

}