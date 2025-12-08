package main

import (
	"fmt"
	"maps"
)

func main(){

	// maps is similar to objects in js or hashmaps

	// syntax: map_name:=make(map[key_type]value_type)
	hashMap:= make(map[string]string)

	hashMap["a"]="apple"
	hashMap["b"]="banana"
	hashMap["c"]="cat"

	fmt.Println("HashMap: ", hashMap)

	// print simgle element:
	fmt.Println("a: ", hashMap["a"])

	// delete a key:
	delete(hashMap, "c");
	fmt.Println("Hashmap: ", hashMap);

	// clear map:
	clear(hashMap)
	fmt.Println("Hashmap: ", hashMap);


	// alternat way to define map:
	// map_name := map[key_type]value_type { key: value }

	newMap := map[string]int {
		"apple": 1,
		"banana": 2,
		"cat": 3,
		"dog": 4,
	}

	fmt.Println("New map: ", newMap);

	// get values from map:
	value,ok :=newMap["apple"]

	fmt.Println("value: ", value);
	fmt.Println("ok: ", ok); // either true or false


	// check if two maps are equal or not:
	map1:= map[string]string {"a": "apple"}
	map2:= map[string]string {"a": "apple"}
	map3:= map[string]string {"b": "banana"}

	fmt.Println("check m1, m2 are equal ? ", maps.Equal(map1, map2))
	fmt.Println("check m2, m3 are equal ? ", maps.Equal(map3, map2))

}
