package main

import (
	"fmt"
)

/*
	NOTES on how Nil maps behave:
	passing a Nil map to delete() nothing happens
	reading a value from Nil map will return the zero value
	add a value to a Nil map the program will cause a panic (new write to a Nil map)
*/

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	fmt.Println("-----------------------------------")
	test3()
}

func test1() {
	// create a map of string to int
	// the key can't be a types that is a slice, a map, a function, or a type that contains any of those 3
	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	// hello in m: 300
	fmt.Println("hello in m:", h)
	// a in m: 0 -- the zero value, which for int is 0
	fmt.Println("a in m:", m["a"])

	// v is the value, ok is a boolean of true if value found, false if not
	if v, ok := m["hello"]; ok {
		// hello in m: 300
		fmt.Println("hello in m:", v)
	}

	// a in m: not found.
	if v, ok := m["a"]; ok {
		fmt.Println("a in m:", v)
	} else {
		fmt.Println("a in m: not found.")
	}

	m["hello"] = 20
	// hello in m: 20
	fmt.Println("hello in m:", m["hello"])
}

func test2() {
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50, // must end with a comma
	}

	// order is randomly retrieved and will not always be a, b, c
	for k, v := range m2 {
		/*
			1st run:
			a 1
			b 2
			c 50

			2nd run:
			c 50
			a 1
			b 2
		*/
		fmt.Println(k, v)
	}

	// b in m2: 2
	fmt.Println("b in m2:", m2["b"])
	delete(m2, "b")
	// b in m2: 0
	fmt.Println("b in m2:", m2["b"])
}

func test3() {
	// declared and initialized
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	// declared, not initialized
	var m3 map[string]int

	// goodbye in m: 0
	fmt.Println("goodbye in m:", m["goodbye"])
	// m3: map[]
	fmt.Println("m3:", m3)
	// maps are referenced types, pointing to the same memory
	m3 = m
	// m3: map[a:1 b:2]
	fmt.Println("m3:", m3)
	m3["goodbye"] = 400
	// goodbye in m3: 400
	fmt.Println("goodbye in m3:", m3["goodbye"])
	// goodbye in m: 400
	fmt.Println("goodbye in m:", m["goodbye"])

	modMap(m)
	// cheese in m: 20
	fmt.Println("cheese in m:", m["cheese"])
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
