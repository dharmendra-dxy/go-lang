package main

import (
	"fmt"
)

func main(){
	// ranges : are used to iterate over data structures

	nums:= []int {1,2,3,4,5,6,7,8,9,10}

	//  FOR LOOP:
	fmt.Println("FOR LOOP: ")

	for i:=0 ; i< len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println("")


	// using RANGES:
	fmt.Println("RANGES")

	for i:= range 5 {
		fmt.Print(nums[i], " ")
	}
	fmt.Println(" ")

	// RANGE to get sum of nums :
	sum:=0
	for _, num := range nums {
		sum = sum + num
	}
	println("The sum is: ", sum)

	// _ used here is the index
	fmt.Println("Index is RANGE:")

	for idx, num:= range nums {
		println("Idx - ", idx, "num - ", num)
	}


	// MAPS AND RANGES
	hashMaps:= map[string]string {
		"fname": "Jhon",
		"lname": "Doe",
		"city": "Noida",
	}

	for key,value:= range hashMaps {
		fmt.Println("key: ", key, "value:", value)
	}


	// STRING and RANGES
	const stringName = "go_lang";

	for idx, c:= range stringName{
		fmt.Println("idx: ", idx, "c: ", c, string(c)); // every char has a unique code
	}

}
