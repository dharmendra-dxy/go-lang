package main

import (
	"fmt"
	"sync"
)

// GO_ROUTINES: used for multy threading in go
// Wait Groups: Mechanism to syncronise the tasks in go in go-routine

func performTask(id int, w *sync.WaitGroup){

	fmt.Println("Performin task for - ", id);

	defer w.Done() // mark waitGroud as done
	// defer is used to execute the process at end

}

func main(){

	// define wait groups:
	var waitGroup sync.WaitGroup


	for i:=0; i<10; i++{
		waitGroup.Add(1) // add in waitGroup
		go performTask(i, &waitGroup) // keyword "go" is used for go-routine
	}

	// main func is also a go-routine
	// put main func in sleep for 2 seconds - so we see go-routines PerformTask()

	// time.Sleep(time.Second*2) // use waitGroup instead

	waitGroup.Wait()
}	
