package main

import "fmt"

func addNumbers(num1 int, num2 int) int {
	return num1+ num2
}

func subtractNumbers(num1 int, num2 int) int {
	return (num1- num2)
}

// in GO: a function can return a multiple value
// syntax: func_name (params..) (return_type1, return_type2, return_type3...)
func getLanguages() (string, string, string) {
	return "go_lang", "java_script", "c++"
}

// in GO: functions can accept func as a paramteres and return them:

// acceppt func as params: 
// func main_function (fn func(num int) int ){} 
// Here: fn is the paramter with type - func(num int) int

// return func as return value:
// func main_function ( )fn func(num int) int{} 
// Here: fn is the returned function with type - func(num int) int

// variadic functions in Go:
// variadic func - can receive n numbers of paramters in a functions

func sumOfNumbers(nums ...int) int {
	// this nums is a slice of int i.e nums:= []int

	sum:=0
	for _,num:= range nums {
		sum= sum+ num
	}
	return sum
}


func main(){

	result1:= addNumbers(5, 10)
	println("Result1: ", result1)

	result2:= subtractNumbers(20,2)
	println("result2: ", result2)


	lang1, lang2, lang3 := getLanguages()
	println("lang1 - ", lang1)
	println("lang2 - ", lang2)
	println("lang3 - ", lang3)

	result4:= sumOfNumbers(2,4,6,8,10)
	fmt.Println("result4: ", result4)

	// pass using slices:
	nums:= []int {3,4,5,6,7,8,9}
	result5:= sumOfNumbers(nums...)
	fmt.Println("result5: ", result5)
}
