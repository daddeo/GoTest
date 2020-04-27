package main

import "fmt"

func main() {
	test1()
	// fmt.Println("-----------------------------------")
	// test2()
	// fmt.Println("-----------------------------------")
	// test3()
	// fmt.Println("-----------------------------------")
}

func test1() {
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++ {
		in <- i
	}

	// means will not have anymore data written to it, but does not wipe out it's contents
	// all values in a buffered channel are still available to be read
	// reading from a closed channel with no more values will immediately return with
	// the zero value for the type specified on the channel.
	//
	// gotcha on close channel:
	// need write to a closed channel or call close a second time on a closed channel (program will panic)
	//
	// when multiple goroutines are writing to a channel, if close the channel, then make sure that all
	// channels are done writing to the channel... and only one goroutines is responsible for closing the
	// channel
	close(in)

	// checking for error (ok) determines if the zero value was returned or a real value of 0 (e.g. int)
	go func() {
		for {
			i, ok := <-in
			if !ok {
				close(out)
				break
			}
			out <- i * 2
		}
	}()

	/*
		outputs:
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
	*/
	for v := range out {
		fmt.Println("v:", v)
	}
}
