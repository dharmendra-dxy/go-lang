package main

import "fmt"

func main(){

	// syntax: var var_name var_type
	// syntax : name = "hello world" -> automatically detects the type
	// short hand synatx: name := "hello world" 

	var name string = "go_lang"
	fmt.Println(name)

	personName := "Dharmendra"
	fmt.Println("Person name is ", personName);

	// int value
	var age = 20
	fmt.Println("Age is ", age);

	// float value
	var price = 500.50
	fmt.Println("Price of the product is ", price)


	// var: when defined, variables can be changed
	// const : when defined, variables can-not be changed

	const teacherName = "Hitesh chaudhary";
	const teacherAge = 35;
	fmt.Println("Teacher Name - ", teacherName)
	fmt.Println("Teacher Age - ", teacherAge)

}
