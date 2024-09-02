package main

import (
	"fmt"
	"sync"
)

// Go that allow for concurrent execution of functions. They are lightweight, efficient threads managed by the Go runtime,
// enabling you to build highly concurrent programs easily.

// 1 - Concurrency vs. Parallelism:
// 2 - Starting a Goroutine: like -> go someFunction()
// 3 - Lightweight and Efficient:
// 4 - Goroutine Scheduling:
// 5 - Communication between Goroutines:
// 6 - Synchronizing Goroutines:

func task(val int, w *sync.WaitGroup) {
	fmt.Println("print Val -> ", val)

	defer w.Done() // complete wait Group function

	//  defer => keyword is used to schedule a function call to be executed after the surrounding function has returned.
	// Deferred function calls are executed in last-in-first-out (LIFO) order, meaning the last deferred function will be executed
	//  first when the function exits.

}

func goroutines_and_waitGroup() {
	var wg sync.WaitGroup // create WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go task(i, &wg) // Go => Goroutines is running paralle work

		// go func(i int) {
		// 	fmt.Println("print ananomiys function Val -> ", i)
		// }(i)
	}
	// time.Sleep(time.Second * 2) // delay to run the code

	wg.Wait() // wait for the function run
}

// WaitGroup => is a synchronization primitive provided by the sync package that allows you to wait for a collection of goroutines to
//finish executing. It is particularly useful when you have multiple goroutines performing tasks concurrently and you need to wait for
// all of them to complete before proceeding with the next steps in your program

// WaitGroup is used to wait for multiple goroutines to complete.
// You use Add to set the number of goroutines to wait for, Done to signal the completion of a goroutine, and Wait to block until all goroutines have finished.
// It’s an essential tool for managing concurrency in Go, ensuring that the main program waits for all necessary work to be done before proceeding.
