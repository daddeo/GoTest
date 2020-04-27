/*
	SELECT
	is most important for managing the done channel pattern

	Go automatically manages memory with garbage collection, it does NOT manage goroutines
	anytime you launch a goroutine you must provide a way for it to exit, if you don't
	resources will be wasted and the go runtime will continue to schedule time for a
	goroutine that will do nothing

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
	test2()
}

func test1() {
	//test1a()
	test1b()
	fmt.Println("-----------------------------------")
	test1c()
	fmt.Println("-----------------------------------")
	test1d()
	fmt.Println("-----------------------------------")
}

/* causes a deadlock since nothing is reading from in as out is waiting */
func test1a() {
	in := make(chan int)
	out := make(chan int, 1)

	out <- 1
	in <- 2
	fmt.Println("write 2 to in")
	v := <-out
	fmt.Println("read", v, "from out")
}

func test1b() {
	in := make(chan int)
	out := make(chan int, 1)

	out <- 1

	// if any case can succeed then that read or write happens and the commands
	// for that case are executed
	//
	// outputs:
	// read 1 from out

	select {
	case in <- 2:
		fmt.Println("write 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}
}

// when multiple cases in a select statement are ready, one is randomly picked to
// avoid certain types of deadlocks and it is executed
func test1c() {
	in := make(chan int, 1)  // buffered
	out := make(chan int, 1) // buffered

	out <- 1

	// if any case can succeed then that read or write happens and the commands
	// for that case are executed
	//
	// outputs:
	// write 2 to in
	// write 2 to in
	// read 1 from out

	select {
	case in <- 2:
		fmt.Println("write 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}
}

// default case is run when none of the cases can read or write to their channels. this will avoid
// blocking reads or writes for channels
// default is NEVER run if any other case in the select can be run
func test1d() {
	in := make(chan int)  // unbuffered
	out := make(chan int) // unbuffered

	// outputs:
	// nothing else works
	// nothing else works
	// nothing else works

	select {
	case in <- 2:
		fmt.Println("write 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	default:
		fmt.Println("nothing else works")
	}
}

func test2() {
	test2a()
	fmt.Println("-----------------------------------")
	test2b()
	fmt.Println("-----------------------------------")
}

func multiples(i int) chan int {
	out := make(chan int)
	curVal := 0
	go func() {
		for {
			out <- curVal * i
			curVal++
		}
	}()
	return out
}

/*
	output:
	v: 0
	v: 2
	v: 4
	v: 6
	v: 8
	v: 10
	v: 12
	v: 14
	v: 16
	v: 18
	v: 20
*/
func test2a() {
	twosChannel := multiples(2)
	// the channel is never closed and open until the program exist
	// .. and so we are wasting resources
	// in a small program this might not be a problem, but in a larger long running program this
	// will eventually swap the Go scheduler
	// so we need to signal to it when it's time to shutdown we need to use a done channel (see test2b)
	for v := range twosChannel {
		if v > 20 {
			break
		}
		fmt.Println("v:", v)
	}

	// do more stuff
}

func multiplesV2(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{}) // channel of empty structs
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				fmt.Println("goroutine shutting down")
				return
			}
		}
	}()
	return out, done
}

/*
	output:
	v: 0
	v: 2
	v: 4
	v: 6
	v: 8
	v: 10
	v: 12
	v: 14
	v: 16
	v: 18
	v: 20
	goroutine shutting down
*/
func test2b() {
	twosChannel, doneChannel := multiplesV2(2)
	// we use the done channel to signal when it's time to shutdown
	for v := range twosChannel {
		if v > 20 {
			break
		}
		fmt.Println("v:", v)
	}
	close(doneChannel)

	// do more stuff, add sleep to fake doing more stuff
	time.Sleep(1 * time.Second)
}
