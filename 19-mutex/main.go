package main

import (
	"fmt"
	"sync"
)

// MUTEX: mutual exclusion: helps to get out of race condition
// race condition : when multiple process are changing the same file, it should be atomic

type post struct {
	views int
	mu sync.Mutex // declared mutex to avoid race condition
}

func (p *post) incrementViews(waitGroup *sync.WaitGroup){
	// defer waitGroup.Done()
	defer func(){
		p.mu.Unlock()
		waitGroup.Done()
	}()


	p.mu.Lock() // lock the mutex
	p.views +=1
	p.mu.Unlock() // unlock once modified.
	// BEST PRACTICE: shift unlock in defer function
}

func main() {
	var waitG sync.WaitGroup

	myPost:= post{
		views: 0,
	}


	// myPost.incrementViews()
	// fmt.Println("VIEWS: ", myPost.views)

	// in real life, the operations are conceurrent i.e. multiple operations can be performed at a same time
	for i:=0; i<100; i++ {
		waitG.Add(1)
		go myPost.incrementViews(&waitG) // go-routine for concurrent operations
	}

	waitG.Wait()
	fmt.Println("VIEWS AFTER MULTIPLE CONCURRENT OPERATIONS: ", myPost.views)

	//myPost.views - different results when run different time
	// REASON: Operations are concurrent. Multiple operations tries to change same value (RACE CONDITION)
	// SOLUTION: 


}
