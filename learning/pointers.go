package main

import "fmt"

func main() {
	fmt.Println("---------- (test1)")
	test1()
	//fmt.Println("---------- (test2)")
	//test2()
	fmt.Println("---------- (test3)")
	test3()
	fmt.Println("---------- (test4)")
	test4()
	fmt.Println("---------- (test5)")
	test5()
}

func test1() {
	a := 10
	b := &a // pointer to a, reference -- b's value is the memory address of a
	c := a
	fmt.Println(a, b, *b, c) // *b, de-references the memory address to get the value

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)
}

func test2() {
	var b *int         // currently set to nil (the absence of a valid value)
	fmt.Println(b, *b) // causes a panic on *b
}

func test3() {
	b := new(int)      // new pointer of int and allocated
	fmt.Println(b, *b) // causes a panic on *b
}

func setTo10(a *int) {
	*a = 10
}

func test4() {
	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}

func setTo10Fail(a *int) {
	// won't work cause the int pointer value is pass-by-refence and cannot be changed here
	a = new(int)
	*a = 10
}

func test5() {
	a := 20
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)
}
