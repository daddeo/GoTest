/*
	Channels
	- a type in Go for transferring data between goroutines
	- one (or more) goroutines write to a channel
	- one (or more) goroutines read back from the same channel (in the order it was written)
	- are intended to be shared between goroutines
	- data on channels is typed
	- by default, channel reads and writes are synchronous
		. when a goroutines writes to a channel it pauses until another goroutine reads from that channel
		. an conversly when a goroutine reads from a channel and no data is present it waits data is available
	- putting a value in a channel is just like passing a value to as function
	- be careful when passing refrence types over a channel
		pass by value will not change the sending data (it's a copy)
		pointer, map, slice pass by reference (a pointer) and
		any changes made to the data in the receiving goroutine will modify the data in the sending goroutine

	- are types in Go, has a zero value (nil)
	- channels are refrence types, like maps, slices, pointers, and functions
	- writing to a closed channel or read/write to a nil channel won't cause a panic and will make your goroutine
		hang forever

	read
		unbuffered: pause until something is written
		buffered: pause if buffer is empty
		nil: hang forever
		closed: return immediately w/ zero value (use comma-ok to see if it is closed)
	write
		unbuffered: pause until something is read
		buffered: pause only if buffer full
		nil: hang forever
		closed: panic
	close
		unbuffered: works
		buffered: works
		nil: panic
		closed: panic
*/

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
	in := make(chan string)
	out := make(chan string)
	go func() {
		name := <-in
		out <- fmt.Sprintf("Hello, " + name)
	}()
	in <- "Bob"
	close(in)
	message := <-out
	// message: Hello, Bob
	fmt.Println("message:", message)
}
