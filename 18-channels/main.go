package main

import (
	"fmt"
	"time"
)

// 1. func to demonstrate: send data from channel
func processNumber(numberChannel chan int) {
	for num := range numberChannel {
		fmt.Println("Processing number - ", num)
		time.Sleep(time.Second * 1)
	}
}

// 2. func to demonstrate: get data to channel
func sumOfNumbers(resultChannel chan int, num1 int, num2 int) {
	sum := num1 + num2
	resultChannel <- sum
}

// 3. func : channels can behave similar like wait_groups:
func processTask(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("Processing......")
}

// 4. func: buffered channel
func emailSender(emailChannel chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChannel {
		fmt.Println("SEND EMAILS TO - ", email)
		time.Sleep(time.Second)
	}
}

 /*  
 INCREASE TYPE SAFETY:
 1. receive only channel : func emailSender(emailChannel <-chan string, done chan bool)
 2. send only channel : func emailSender(emailChannel chan<- string , done chan bool)
 */

func main() {

	// CHANNEL : It enables to send data from one go-routine to another

	// ====================== 1. BASICS ==========================

	// create a channel:
	// messageChannel := make(chan string)

	// send data to a channel:
	// messageChannel <- "random_message" // channels are blocking -

	// get data from a channel:
	// data:= <-messageChannel
	// fmt.Println("Data: ", data)

	/*
		ERROR: when running the file
		fatal error: all goroutines are asleep - deadlock!
		reason : messageChannel <- "random_message": it is a blocking event
		sending | receiving events are "BLOCKING"
	*/

	// ============= 2. SEND DATA FROM A CHANNEL =============

	// numberChannel := make(chan int)
	// go processNumber(numberChannel) // go-routine - it is a non-blocking event

	// generally, we use as a queue. We send multiple data
	// numberChannel <- 5

	// infinite loop:
	// for {
	// 	numberChannel <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 2)

	// ===================== 3. GET DATA IN CHANNEL ====================

	// resultChannel:= make(chan int)
	// go sumOfNumbers(resultChannel, 4,5)

	// result:= <-resultChannel
	// fmt.Println("ResultChannel - ", result)

	// =================== 4. CHANNEL behaviour is similar to wait_groups =====================

	// doneChannel:= make(chan bool)
	// go processTask(doneChannel)

	// <- doneChannel // block for now - so we can see the output

	// ======= 5. BUFFERED CHANNEL - sending | receving operations are "NOT BLOCKING" ===========

	// emailChannel := make(chan string, 100) // 100 is the size of the channel
	// doneChannel := make(chan bool)

	// // emailChannel <- "1@emxale.com"
	// // emailChannel <- "2@emxale.com"

	// go emailSender(emailChannel, doneChannel)

	// for i := range 20 {
	// 	emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// // fmt.Println("emailChannel: ", <-emailChannel);
	// // fmt.Println("emailChannel: ", <-emailChannel);
	// fmt.Println("DONE SENDING THE CHANNEL")
	// close(emailChannel) // close the emailChannel
	// <-doneChannel // block

	// ========= 6. Working on multiple channel =================
	channle1 := make(chan int)
	channle2 := make(chan string)

	go func() {
		channle1 <- 10
	}()

	go func() {
		channle2 <- "hello_duniya"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channle1value := <-channle1:
			fmt.Println("RECEIVED FROM CHANNEL_1 : ", channle1value)
		case channle2value := <-channle2:
			fmt.Println("RECEIVED FROM CHANNEL_2 : ", channle2value)
		}
	}

}
