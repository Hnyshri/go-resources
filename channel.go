package main

import (
	"fmt"
	"time"
)

// a channel is a powerful concurrency primitive that allows goroutines to communicate with each other by sending and receiving values.
//  Channels are typically used to synchronize operations or to share data between goroutines.

// if you have single goroutine then use channel method
// if you have multiple goroutine then use waitGroup method

// In Go, channels can be either unbuffered or buffered.
// The difference between them lies in how they handle the sending and receiving of data.

// 1- Unbuffered Channels -> slow channel = Ex. FIFO
// An unbuffered channel requires both the sender and the receiver to be ready at the same time. The sender will block until the receiver is ready to receive the data, and the receiver will block until there is data available to receive.

// 2 - Buffered Channels -> faster channel = Ex. Bulk Messaging Queue
// A buffered channel has a capacity defined when it's created. The sender can send data to the channel until the buffer is full, without waiting for a receiver. The receiver can then receive data from the channel at its own pace.

func channelFunction() {

	// ==========================================important Part=================================================================

	bufferedChannelFoeQueue := make(chan string)  // buffered Channel
	taskIsDoneforBufferChannel := make(chan bool) // unbuffered Channel

	go sendingBulkEmailQueue(bufferedChannelFoeQueue, taskIsDoneforBufferChannel)

	for i := 0; i < 1; i++ {
		bufferedChannelFoeQueue <- fmt.Sprintf("%d@gmail.com", i)
	}

	// bufferedChannelFoeQueue <- "1@gmail.com"
	// bufferedChannelFoeQueue <- "2@gmail.com"

	fmt.Println("All email send in the Queues")

	close(bufferedChannelFoeQueue) // close the channel to handle the deadlock condition

	<-taskIsDoneforBufferChannel

	// ==========================================Receieving the multiple channel Data=================================

	channel1 := make(chan int)
	channel2 := make(chan string)

	go func() {
		channel1 <- 102
	}()
	go func() {
		channel2 <- "Gupta"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channel1Value := <-channel1:
			fmt.Println("Received Data from Channel 1", channel1Value)
		case channel2Value := <-channel2:
			fmt.Println("Received Data from Channel 2", channel2Value)
		}
	}

	// ===========================================================================================================

	isTaskDone := make(chan bool) // it is a unbuffered channel
	go isTaskDoneFun(isTaskDone)
	// <-isTaskDone
	fmt.Println("goroutine synchromizer call", <-isTaskDone)

	// ===========================================================================================================

	numberChannel := make(chan int) // it is a unbuffered channel

	// go sendingChannelData(numberChannel) // send channel to goroutine
	// for {
	// 	numberChannel <- rand.Intn(100) //  Sending the data to Channel so no need to implement the time sleep
	// }

	// =============================-------------------==============================

	go sumForRecieving(numberChannel, 2, 5)
	res := <-numberChannel // reciveing the blocking so no need to implement the time sleep
	fmt.Println(res)

	// numberChannel <- 3 //  Sending the data to Channel

	// time.Sleep(time.Second * 2)

	// =============================-------------------==============================

	// makeDeadLockChannel := make(chan string) // it is a unbuffered channel
	// makeDeadLockChannel <- "shriyansh"
	// // sending and receiveing is blocking. Sending the data to Channel
	// // It will generate the deadlock -> fatal error: all goroutines are asleep - deadlock!

	// messageRecived := <-makeDeadLockChannel // recieving the data from channel

	// fmt.Println(messageRecived)
}

func sendingBulkEmailQueue(emailChan <-chan string, taskIsDone chan<- bool) {
	// <-chan means receive-only channel emailChan
	// chan<- means receive from send-only channel

	defer func() {
		taskIsDone <- true
	}()

	// <-taskIsDone // giving error
	// emailChan <- "dd" // giving error

	for email := range emailChan {
		fmt.Println("Email Send from Queue", email)
		time.Sleep(time.Second) // Default 1 second
	}
}

func sendingChannelData(num chan int) {
	for number := range num {
		fmt.Println("Fetch or recieve the value from channel", number) // no need to write <- number when use loop
		time.Sleep(time.Second)
	}
}

func sumForRecieving(getChann chan int, num1 int, num2 int) {
	sum := num1 * num2
	getChann <- sum // sending in blocking so no need to implement the time sleep
}

// goroutine synchromizer
func isTaskDoneFun(status chan bool) {

	defer func() { // defer is use for cleaning. defer will call after excuting the function
		status <- true
	}()

	fmt.Println("Task has been started...")
}
