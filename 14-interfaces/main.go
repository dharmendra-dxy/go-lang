package main

import "fmt"

// interface syntax: type name (generally as - namer)  interface{}
type paymenter interface {
	pay(amount float32) // pay is signature method that every gateway must need to have
}

type payment struct{
	gateway paymenter // now gateway can we razorpay|stripe
}

func (p payment) makePayment (amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}
func (r razorpay) pay (amout float32){
	fmt.Println("Payment made using razorpay ", amout)
}

type stripe struct {}
func (s stripe) pay (amout float32){
	fmt.Println("Payment made using stripe ", amout)
}

func main(){

	razorpayPaymentGatweay:= razorpay{}

	newPayment := payment{
		gateway: razorpayPaymentGatweay,
	}

	newPayment.makePayment(600)
}
