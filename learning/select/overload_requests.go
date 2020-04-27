package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1,000 --> total time: 1.0109954s
	// 10,000 --> total time: 1.1000947s
	// 100,000 --> total time: 2.1759983s
	runIterations(1) // 10000

	/*
		USING BACKPRESSURE makes it harder for you process to be overwhelmed by requests, simply
		ignore the ones you don't have capacity to handle.

		10,000 workers and 10,000 iterations
		total processed: 10000
		total time: 1.1258251s

		10,000 workers and 100,000 iterations
		total processed: 10000
		total time: 1.1640024s

		20,000 workers and 100,000 iterations
		total processed: 20000
		total time: 1.3540045s
	*/
	runLimiter(20000, 100000)
}

func runIterations(iterations int) {
	var wg sync.WaitGroup
	totalStart := time.Now()
	// launch 10,000 goroutines
	for i := 0; i < iterations; i++ {
		start := time.Now()
		wg.Add(1)
		go func(in int) {
			time.Sleep(1 * time.Second) // sleep for a second
			out := 2 * in
			fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("total time:", time.Now().Sub(totalStart))
}

func runLimiter(workers int, iterations int) {
	//workers := iterations
	// create a pool of workers using a buffered channel of functions that take in an int and return an int
	pool := make(chan func(int) int, workers)
	// a pool of functions
	for i := 0; i < workers; i++ {
		pool <- func(in int) int {
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < iterations; i++ {
		start := time.Now()
		select {
		// try and get a worker from the pool, does work and then returns the worker to the pool
		case f := <-pool:
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)
		default:
			fmt.Println("rejecting request", i, "too busy")
		}
	}

	wg.Wait()

	fmt.Println("total processed:", count)
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
