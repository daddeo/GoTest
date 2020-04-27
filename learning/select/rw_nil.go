/*
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	test1()
	// fmt.Println("-----------------------------------")
	// test2()
}

/*
	output:
	from in, result is 0
	from in2, result is 102
	from in, result is 2
	from in2, result is 103
	from in, result is 4
	from in2, result is 104
	from in, result is 6
	from in2, result is 105
	from in, result is 8
	from in2, result is 106
	from in2, result is 107
	from in, result is 10
	from in2, result is 108
	from in, result is 12
	from in2, result is 109
	from in2, result is 110
	from in2, result is 111
	from in, result is 14
	from in, result is 16
	from in, result is 18
*/
func test1() {
	in := make(chan int)
	in2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			in2 <- i
		}
		close(in2)
		wg.Done()
	}()

	go func() {
		count := 0
		// run until count == 2
		for count < 2 {
			select {
			case i, ok := <-in:
				if !ok {
					count++
					// do this cause a read from a closed channel returns a value immediately, but
					// a read from a nil channel never returns and since we've already closed the
					// channel we want to stop trying to read from them as there is no more data
					// coming in. effectively disables the case in the select and when both are set
					// to nil, count is incremented to do and effectively causes us to exit the for
					// loop and exit the goroutine
					in = nil
					continue
				}
				fmt.Println("from in, result is", i*2)
			case i, ok := <-in2:
				if !ok {
					count++
					// do this cause a read from a closed channel returns a value immediately, but
					// a read from a nil channel never returns and since we've already closed the
					// channel we want to stop trying to read from them as there is no more data
					// coming in. effectively disables the case in the select and when both are set
					// to nil, count is incremented to do and effectively causes us to exit the for
					// loop and exit the goroutine
					in2 = nil
					continue
				}
				fmt.Println("from in2, result is", i+2)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
