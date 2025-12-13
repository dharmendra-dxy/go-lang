package main

import "fmt"

// by default: pass by value
// the value is passed, the original number remains un-effected
func getNumber (num int) {
	num=10
	fmt.Println("FUNCITON : getNumber : ", num)
}

// pass by references:
func getNumerChanged(num *int){
	*num=100
	fmt.Println("FUNCTION: getNumberChanged: ", num);
}

func main(){

	// POINTER: pointer defines the memory location of the variable
	number:=5
	getNumber(number)
	fmt.Println("MAIN : Number: ", number) // the value does not get changes, remains 5

	// REASON : Pass by value, Pass by reference

	newNumber:=10

	fmt.Println("Memory adress of newnumber: ", &newNumber)

	getNumerChanged(&newNumber)
	fmt.Println("MAIN: newNumber: ", newNumber)
}
