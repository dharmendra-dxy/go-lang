package main

import "fmt"

// enums : enumarated types
type OrderStatus int

const (
	Received OrderStatus = iota // iota is auto incremental - 0
	Confirmed // 1
	Prepared
	Delivered
)

// for orderStatus string:
type NewOrderStatus string
const (
	S_Received NewOrderStatus = "received" 
	S_Confirmed NewOrderStatus = "confirmed"
	S_Prepared NewOrderStatus = "prepared"
	S_Delivered NewOrderStatus = "delivered"
)

func changeOrderStatus (status OrderStatus) {
	fmt.Println("order status changed to - ", status)
}

func main(){

	changeOrderStatus(Delivered)

}