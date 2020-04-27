package main

import (
	"fmt"
	"sync"
)

func runHello(name string) {
	fmt.Println("Hello to", name, "from a goroutine.")
}

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	fmt.Println("-----------------------------------")
	test3()
	// fmt.Println("-----------------------------------")
}

/*
	output:

	Hello to noob from a goroutine.
*/
func test1() {
	var wg sync.WaitGroup
	wg.Add(1)
	// 'go' runs a closure (aka anonlymous function)
	go func(name string) {
		runHello(name)
		wg.Done()
	}("noob")

	// pauses the main goroutine until the count of running goroutines, tracked by wg, is zero
	wg.Wait()
}

/*
	output:

	2
	7
	7
	8
	10
	10
	10
	10
	2
	10
*/
func test2() {
	var wg sync.WaitGroup

	// in this case the value of i is shared by all goroutines with most not getting the value until after the for
	// loop has finished, hence the value of 10
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 'go' runs a closure (aka anonlymous function)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	// pauses the main goroutine until the count of running goroutines, tracked by wg, is zero
	wg.Wait()
}

/*
	output:

	4
	2
	6
	5
	0
	7
	8
	9
	1
	3
*/
func test3() {
	var wg sync.WaitGroup

	// pass the value of i into the closure to fix the problem in test2()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 'go' runs a closure (aka anonlymous function)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}

	// pauses the main goroutine until the count of running goroutines, tracked by wg, is zero
	wg.Wait()
}
