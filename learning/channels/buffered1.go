/*
	use a buffered channel when you do want a channel to wait of read or write
*/

package main

import "fmt"

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	// fmt.Println("-----------------------------------")
	// test3()
	// fmt.Println("-----------------------------------")
}

func test1() {
	// channel with room for 10 items of size int
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out <- localI * 2
		}(i)
	}
	// int slice
	var result []int
	for i := 0; i < 10; i++ {
		val := <-out
		result = append(result, val)
	}

	// result: [0 2 6 8 10 18 12 14 16 4]
	fmt.Println("result:", result)
}

/*
	when communicating between goroutines, you need to make sure that atleast one
	goroutine is not blocked reading from or writing to a channel
*/
func test2() {
	// this test will deadlock!!!
	//test2a()

	// BUT
	// adding size 2 to make channel will solve in this case... not a viable long term solution,
	// since there could easily be a future case when panic will happen again.
	test2b()

	// instead, look for ways to fix the problem. In this example doing a read before the second
	// write solves the problem.
	test2c()
}

func test2a() {
	in := make(chan int)
	out := make(chan int)

	// loops forever reading an int from the in channel and writing the double to the out channel
	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	// write twice to the in channel
	in <- 1
	in <- 2
	// read twice from the out channel
	o1 := <-out
	o2 := <-out

	// o1: 2 , o2: 4
	fmt.Println("o1:", o1, ", o2:", o2)
}

func test2b() {
	in := make(chan int, 2)
	out := make(chan int, 2)

	// loops forever reading an int from the in channel and writing the double to the out channel
	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	// write twice to the in channel
	in <- 1
	in <- 2
	// read twice from the out channel
	o1 := <-out
	o2 := <-out

	// o1: 2 , o2: 4
	fmt.Println("o1:", o1, ", o2:", o2)
}

func test2c() {
	in := make(chan int)
	out := make(chan int)

	// loops forever reading an int from the in channel and writing the double to the out channel
	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	// write to the in channel
	in <- 1
	// read from the out channel
	o1 := <-out
	// write to the in channel
	in <- 2
	// read from the out channel
	o2 := <-out

	// o1: 2 , o2: 4
	fmt.Println("o1:", o1, ", o2:", o2)
}
