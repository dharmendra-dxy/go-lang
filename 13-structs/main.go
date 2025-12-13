package main

import (
	"fmt"
	"time"
)

// in other programming languages: we use classes and object
// in GO we use: struct

// eg struct - syntax: type struct_name struct {}

type customer struct {
	name string
	age int
}

type orders struct{
	id string
	name string
	amount float32
	status string
	createdAt time.Time

	customer //struct embedding
}



// alternative of constructor in GO:
// constructors are used to create instances of classes
func newOrders(id string, name string, amount float32, status string) *orders{
	myOrder := orders{
		id: id,
		name: name,
		amount: amount,
		status: status,
	}

	return &myOrder // we return the pointer of the instance
}

// methods/function can be attached with structs
// receiver type 
// use PASS BY REFERENCE in struct when need to update any value
func (o *orders) changeStaus(status string){
	o.status = status

	// *o.status = status - no need to use de-reference 
}

func (o orders) getOrderAmount() float32{
	return o.amount
}


func main(){

	// in instance - if we don't set the value, the default value is zero value (0, nil, "", false)

	// instance of order struct:
	myOrder := orders{
		id: "1",
		name: "Pizza",
		amount: 132.34,
		status: "pending",
	}
	myOrder.createdAt = time.Now();
	fmt.Println("myOrders: ", myOrder)


	// changeStatus:
	myOrder.changeStaus("paid")
	fmt.Println("myOrders: ", myOrder)

	// print orderAmount:
	fmt.Println("orderAmount: ", myOrder.getOrderAmount())

	// use constructor:
	myOrder2 := newOrders("2", "Cake", 199.0, "pending")
	fmt.Println("myOrder2: ", myOrder2);
	fmt.Println("myOrder2 Price: ", myOrder2.getOrderAmount())


	// struct embedding usage:
	myCustomer:= customer{
		name: "Jhon Doe",
		age: 29,
	}

	myOrder3 := orders {
		id: "3",
		name: "Panner Garlic",
		amount: 399,
		status: "pending",
		customer: myCustomer,
	} 

	fmt.Println("myOrder3: ", myOrder3)
	fmt.Println("order3 customer name: ", myOrder3.customer.name)

}
