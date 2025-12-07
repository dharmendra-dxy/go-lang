package main

import "fmt"

func main(){

	// in go for loops we use - "for"

	// loop to print 1 to 10:
	fmt.Println("For loop")
	var count=10;
	for i := 0; i < count; i++ {
		fmt.Print(i , " ");
	}
	fmt.Println("")


	// range can be alos used in go:
	fmt.Println("RANGE: ")
	for k:= range 5 {
		fmt.Print(k, " ");
	} 
	fmt.Println("")


	// while loop:
	fmt.Println("While loop")

	var j=1;
	for j<=5 {
		fmt.Print(j, " ")
		j = j+1
	}

}
