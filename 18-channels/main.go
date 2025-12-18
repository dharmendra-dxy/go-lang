package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func to demonstrate: send a data from go routine
func processNumber(numberChannel chan int) {
	for num := range numberChannel {
		fmt.Println("Processing number - ", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// CHANNEL : It enables to send data from one go-routine to another

	// create a channel:
	// messageChannel := make(chan string)

	// send data to a channel:
	// messageChannel <- "random_message" // channels are blocking

	// get data from a channel:
	// data:= <-messageChannel
	// fmt.Println("Data: ", data)

	/*
		ERROR: when running the file
		fatal error: all goroutines are asleep - deadlock!
		reason : messageChannel <- "random_message": it is a blocking event
	*/

	numberChannel := make(chan int)
	go processNumber(numberChannel) // go-routine - it is a non-blocking event

	// generally, we use as a queue. We send multiple data
	// numberChannel <- 5

	// infinite loop:
	for {
		numberChannel <- rand.Intn(100)
	}

	// time.Sleep(time.Second * 2)
}
