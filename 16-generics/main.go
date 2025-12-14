package main

import (
	"fmt"
)

// we are repeating our code for different data_types - use generics

func printSlice(items []int) {
	fmt.Println("Print Slice: ")
	for _,item:= range items {
		fmt.Print(item, " ");
	}
	fmt.Println("")
}

func printStringSlice(items []string) {
	fmt.Println("Print String Slice: ")
	for _,item:= range items {
		fmt.Print(item, " ");
	}
	fmt.Println("")
}

// GENERIC FUNCTION
func printGenericSlice[T int|string ](items []T){
	fmt.Println("Print generic Slice: ")
	for _,item:= range items {
		fmt.Print(item, " ");
	}
	fmt.Println("")
}

// GENERICS on struct
type stack[T int| string] struct {
	element []T
}

func main(){

	nums:= []int {8,3,59,32,43,2}
	names:= []string {"Ayush", "Aniket", "Pulak"}

	printSlice(nums)
	printStringSlice(names)

	printGenericSlice(names)
	printGenericSlice(nums)

	// struct
	myStack := stack[int]{
		element: []int {3,5,7,7,3,9,0,2},
	}

	myStack2 := stack[string]{
		element: []string {"go_lang", "ts", "js", "c"},
	}

	fmt.Println("myStack element: ", myStack.element)
	fmt.Println("myStack2 element: ", myStack2.element)
}
