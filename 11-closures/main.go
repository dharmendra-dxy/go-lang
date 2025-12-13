package main

func counter()  func()int {
	var count int=0

	return func() int {
		count = count+1;
		return count;
	}
}

func main(){
	increaseCount:= counter()
	println(increaseCount())
	println(increaseCount())
	println(increaseCount())
	println(increaseCount())
}
