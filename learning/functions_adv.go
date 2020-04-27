package main

import "fmt"

func main() {
	myAddOne_1 := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne_1(1))

	// anonymous function declaration
	myAddOne_2 := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne_2(1))

	printOperation(1, addOne)
	printOperation(1, addTwo)

	local_in_func()

	addOne_3 := makeAdder(1) // a closure instance assigned to a local variable
	addTwo_3 := makeAdder(2)

	fmt.Println(addOne_3(1))
	fmt.Println(addTwo_3(1))

	addOne_4 := makeAdder(1) // a closure instance assigned to a local variable
	doubleAddOne_4 := makeDoubler(addOne_4)

	fmt.Println(addOne_4(1))
	fmt.Println(doubleAddOne_4(1))
}

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func local_in_func() {
	b := 2
	myAddOne := func(a int) int {
		return a + b
	}
	fmt.Println(myAddOne(1))

	// a local variable can be altered by a local anonymous funcion
	c := b
	myAddOne_2 := func(a int) {
		c = a + c
	}
	myAddOne_2(1)
	fmt.Println(c)
}

// uses a closure
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	// our closure, the embedded anonymous function
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}
