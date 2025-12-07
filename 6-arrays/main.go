package main

import "fmt"

func main(){

	// syntax: var var_name [array_size]arr_type

	// declare array:
	var nums [4]int

	// add elements:
	nums[0] = 101
	nums[1] = 102
	nums[2] = 103
	nums[3] = 104

	// array length:
	println("size: ", len(nums))

	println("nums[0] ", nums[0])
	println("nums[1] ", nums[1])
	println("nums[2] ", nums[2])
	println("nums[3] ", nums[3])

	fmt.Println("nums ", nums)

	// declare in single line:
	numbers := [3]int {203, 403, 323}
	fmt.Println("Numbers: ", numbers)

	// 2d arrays
	twoDimension := [2][2]int {{3,4}, {5,6}}
	fmt.Println("twoDimensions : ", twoDimension)

	// in go lang, we use slice instead of array.
	// array - is fixed size, while slices - variable and dynamic sizes

}