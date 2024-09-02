package main

import (
	"fmt"
	"sync"
)

// Mutex => (short for "mutual exclusion") is a synchronization primitive used to protect shared resources from being accesse
// concurrently by multiple goroutines. A Mutex ensures that only one goroutine can access the critical section of code that it
// guards at a time, preventing race conditions.

// Race condition => when two or more goroutines access shared data concurrently, and at least one of them modifies the data.
// If the access and modification of shared data are not properly synchronized, the program's outcome becomes unreliable and may vary
// from run to run.

type post struct {
	view int
	// single resource ko multiple process ko ek sath modify karte hai to ek process dusre process ko override krta hai -
	// So resolve krne Mutex use krte hai then Lock the process on modification

	mutex sync.Mutex // handle the race condition
}

func (p *post) counter(wg *sync.WaitGroup) {

	defer func() {
		p.mutex.Unlock() // best practice
		wg.Done()
	}()

	p.mutex.Lock()
	p.view += 1
	// p.mutex.Unlock()
}
func mutexFunction() {
	var wg sync.WaitGroup

	intialCount := post{view: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go intialCount.counter(&wg) // concurrently goroutine or mupltiple process will excuted
	}
	wg.Wait()
	fmt.Println("Post Coun is ", intialCount.view)
}
