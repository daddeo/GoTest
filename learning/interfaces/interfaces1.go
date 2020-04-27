package main

import (
	"fmt"
	"net/http"
)

/*
	with interfaces:

	there is not explicit declaration of interface implementation
	as long as your type implements the rights methods it automatically meets the interface;
	sorta like duck type in python, ruby or javascript it is assumed it is save to call.

	implementors do not derive or specify they implement and interface, but simlpy implmenting
	the function of the interface (via prototype) they "implement" the interface.
*/

type bound struct {
	lower int
	upper int
}

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2(false)
	fmt.Println("-----------------------------------")
	test3()
	fmt.Println("-----------------------------------")

	// input number { 10 } is between { 5 and 20 } and divisible by { 5 }: true
	limits := bound{lower: 5, upper: 20}
	test4(10, limits, 5)

	// input number { 10 } is between { 10 and 20 } and divisible by { 5 }: false
	limits = bound{lower: 10, upper: 20}
	test4(10, limits, 5)

	// input number { 10 } is between { 5 and 20 } and divisible by { 3 }: false
	limits = bound{lower: 5, upper: 20}
	test4(10, limits, 3)
}

func test1() {
	n := 10
	lower := 5
	upper := 20
	divis := 5
	result := runTests(n, []tester{
		rangeTest{min: lower, max: upper},
		divTest(divis),
	})
	// input number { 10 } is between { 5 and 20 } and divisible by { 5 }: true
	fmt.Println("input number {", n, "} is between {", lower, "and", upper, "} and divisible by {", divis, "}:", result)
}

func test2(showPanic bool) {
	// the empty interface
	// since the interface has no methods, any types in Go match it, even types that don't define methods at all (e.g. built-in types)
	// MORE: like the void type in C or Object type in java; a way to say this could be anything
	var i interface{}
	i = "Hello"
	// asserts the concrete type behind i is string and assigns to j; this is a type assertion
	// MORE: when trying to keep type assertion and type conversion straight use:
	// type conversion: changes the type (e.g. int -> float64); can be used on any type
	// type assertion: you are revealing the underlying type and stripping away the interface that wraps it; only can be done on an interface
	j := i.(string)
	// ok assigned true if assertion worked, false if failed
	// on fail, the value (k) will be assigned the zero value for that type (e.g. int: 0, string: "", float64: 0.0; pointer: Nil)
	k, ok := i.(int)
	// j: Hello , k: 0 , ok: false
	fmt.Println("j:", j, ", k:", k, ", ok:", ok)

	if showPanic {
		// invalid type assertion without the ok will panic
		m := i.(int)
		// panic: interface conversion: interface {} is string, not int
		fmt.Println("m:", m)
	}
}

func doStuff(i interface{}) {
	// a type switch, (type) is a convention
	// since we aren't using the name of a concreate type and ...
	// by convention the variable on the left hand side of := reuses the interface variable's name on the right side
	switch i := i.(type) {
	case int:
		fmt.Println("Double i is", i+i)
	case string:
		fmt.Println("i is", len(i), "characters long")
	default:
		fmt.Println("I don't know what to do with this.")
	}
}

func test3() {
	// Double i is 20
	doStuff(10)
	// i is 5 characters long
	doStuff("Hello")
	// I don't know what to do with this.
	doStuff(true)
}

// we know that Go allows us to declare our own concrete types based on any primitive types, not just structs
// interfaces can match methods on any type not just structs
// you can make a function implement an interface by defining a function type and a method on the function type

type testerEx interface {
	test(int) bool
}

func runTestsEx(i int, tests []testerEx) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type testerExFunc func(int) bool

func (tf testerExFunc) test(i int) bool {
	return tf(i)
}

func test4(n int, b bound, mod int) {
	// n := 10
	// lower := 5
	// upper := 20
	//divis := 5

	// testerExFunc using anonymous functions
	result := runTestsEx(n, []testerEx{
		testerExFunc(func(i int) bool {
			return i%mod == 0
		}),
		testerExFunc(func(i int) bool {
			return i < b.upper
		}),
		testerExFunc(func(i int) bool {
			return i > b.lower
		}),
	})
	fmt.Println("input number {", n, "} is between {", b.lower, "and", b.upper, "} and divisible by {", mod, "}:", result)
}

/*
	converting functions to interface implementations is increbily useful

	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})

	breaks down into:
	type HandleFunc func(http.ResponseWriter, *http.Request)

	func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		f(w,r)
	}

	type Handler interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}

	func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}

	func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		mux.Handle(pattern, HandlerFunc(handler))
	}
*/

func test5() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
