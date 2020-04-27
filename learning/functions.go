package main

import "fmt"

func main() {
	a := addNumbers(2, 3)
	fmt.Println(a)

	a = addNumbers(4, 10)
	fmt.Println(a)

	a = addNumbers(100, -100)
	fmt.Println(a)

	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20)

	v := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(v, arr, s)
	fmt.Println("in main:", v, arr, s)
}

func addNumbers(x int, y int) int {
	return x + y
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

// function parameters are pass by value not by reference
func doubleFail(v int, arr [2]int, s string) {
	v *= 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s += s
	fmt.Println("in doubleFail:", v, arr, s)
}
