package main

import (
	"fmt"
	"time"
)


func main(){

	// if-else and if-else-if statement:
	println("====== IF ELSE ======")

	const age = 11;

	if age>=18 {
		fmt.Println("You are an adult");
	} else if age>=12 {
		fmt.Println("You are a kid");
	} else {
		fmt.Println("You are child");
	}

	// switch statement in go
	println("======== SWITCH CASE =======")

	// simple switch
	const day = "Friday";

	switch day {
		case "Monday": 
			fmt.Println("Today is Monday")
			// break : no need to mention break in golang
		case "Tuesday": 
			fmt.Println("Today is Tuesday")
		case "Wednesday": 
			fmt.Println("Today is Wednesday")
		case "Thursday": 
			fmt.Println("Today is Monday")
		default: 
			fmt.Println("Enjoy, it is weekend")

			// default is optional
	}


	// Multiple condition switch in go:
	println("====== MULTIPLE CONDITION SWITCH =======")

	switch time.Now().Weekday() {

		// multiple conditions
	case time.Saturday , time.Sunday :
		fmt.Println("Hurray, weekend")
	
		default : 
		fmt.Println("Its workday")
	}


	// TYPE SWITCH
	// we can pass function in switch case
	fmt.Println("======== TYPE SWITCH =======")

	getDetail := func(i interface {}) {
		switch t:= i.(type) {
		case int : 
			fmt.Println("It is interger")

		case string : 
			fmt.Println("It is string")

		case bool : 
			fmt.Println("It is a boolean")

		default: 
			fmt.Println("It is a other type ", t)
		}
	}

	getDetail(1);
	getDetail("go lang");
	getDetail(true);


}
