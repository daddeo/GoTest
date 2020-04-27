/*
	Concurrency in Go does not use threading like other languages and is built on CSP (Communicating Sequential Processes)

	Goroutines are lightweight processes managed by the Go runtime (not threads)
	Threads are managed by the operating system.
	Goroutines are managed by the Go runtime scheduler.
	The Go runtime creates a number of threads at startup
	The Go runtime schedules goroutines across threads automatically
	Goroutines are faster to create, more efficient with memory, and faster to switch
	Goroutines creating is faster than thread creation cause you aren't creating an operating system level resource
	Goroutines stack sizes are smaller than thread stack sizes and can grow as needed
	Goroutines are faster to switch than threads
	Can have tens of thousands in a single process.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func runMe() {
	fmt.Println("runMe says Hi")
}

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	// fmt.Println("-----------------------------------")
	// test3()
	// fmt.Println("-----------------------------------")
}

func test1() {
	go runMe()
	// never use in real program!!!
	time.Sleep(1 * time.Second)
}

func test2() {
	var wg sync.WaitGroup
	wg.Add(1)
	// 'go' runs a closure (aka anonlymous function)
	go func() {
		runMe()
		wg.Done()
	}()

	// pauses the main goroutine until the count of running goroutines, tracked by wg, is zero
	wg.Wait()
}
