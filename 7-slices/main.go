package main

import "fmt"

func main() {

	// slice - dynamic array - no fixed size array
	// most used construct in golang

	// syntax: var var_name []type

	// uninitialized array:
	var nums []int

	fmt.Println("nums: ", nums);
	fmt.Println(nums==nil) // true beacuase, nums is un-initialized and null
	fmt.Println("length: ", len(nums));


	// define a new array with value:

	// syntax: make([]type, initial_size, capacity)
	// capacity is the maximum length of the array, by default equals initial_size
	
	var nums2 = make([]int , 2, 10);
	fmt.Println("nums_2 size: ", len(nums2))
	fmt.Println("capacity nums2: ", cap(nums2))

	// add element:
	nums2 = append(nums2, 34);
	nums2 = append(nums2, 35);
	nums2 = append(nums2, 36);

	fmt.Println("num2: ", nums2);

	nums2 = append(nums2, 51);
	nums2 = append(nums2, 52);
	nums2 = append(nums2, 53);
	nums2 = append(nums2, 54);
	nums2 = append(nums2, 55);
	nums2 = append(nums2, 56);
	nums2 = append(nums2, 57);

	fmt.Println("nums2: ", nums2);
	fmt.Println("New Size nums2: ", len(nums2))
	fmt.Println("New capacity num2: ", cap(nums2))
	// capacity increase automatically when sizes increases


	// Alternate way of defining slices:
	num3:= []string {}

	num3 = append(num3, "aniket")
	num3 = append(num3, "yash")
	num3 = append(num3, "pulak")

	fmt.Println("array3: ", num3);
	fmt.Println("Size array3: ", len(num3))
	fmt.Println("capacity array3: ", cap(num3))


	/* operations that can be performed with slices: 
		1. copy: copy(final_arr, base_arr)
		2. slice: printLn(array[start_idx, end_idx])
	*/
	

}
