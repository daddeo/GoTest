package main

import "fmt"

// Foo is a test struct, visible outside of package wnd containing A (visible) and b (not).
type Foo struct {
	A int    // visible outside of package
	B string // not visible outside of package
}

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	fmt.Println("-----------------------------------")
	test3()
	// fmt.Println("-----------------------------------")
	// test4()
}

// a method declaration -- (f Foo); like self in python and this in java
// (f Foo) ==> a value receiver; when not modifying values in the struct

// String does something
func (f Foo) String() string {
	return fmt.Sprintf("A: %d, B: %s", f.A, f.B)
}

// (f *Foo) ==> a reference receiver; when modifying values in the struct

// Double does something
func (f *Foo) Double() {
	f.A = f.A * 2
}

func test1() {
	f := Foo{
		A: 10,
		B: "Hello",
	}

	// A: 10, B: Hello
	fmt.Println(f.String())
	f.Double()
	// A: 20, B: Hello
	fmt.Println(f.String())
}

func (f Foo) fieldCount() int {
	return 2
}

// StringEx does something
func (f Foo) StringEx() string {
	return fmt.Sprintf("%d fields --> A: %d, B: %s", f.fieldCount(), f.A, f.B)
}

// Bar is a test struct
type Bar struct {
	C bool
	Foo
}

func (b Bar) String() string {
	return fmt.Sprintf("b.Foo: %s, b.C: %v", b.Foo.StringEx(), b.C)
}

func (b Bar) fieldCount() int {
	return 3
}

func test2() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	b := Bar{
		C:   true,
		Foo: f,
	}

	// fieldCount is wrong, because there is no method overriding or virtual method dispatch in Go
	// so Foo's fieldCount is called

	// b: b.Foo: 2 fields --> A: 10, B: Hello, b.C: true
	fmt.Println("b:", b.String())
	b.Double()
	// b: b.Foo: 2 fields --> A: 20, B: Hello, b.C: true
	fmt.Println("b:", b.String())
}

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

func test3() {
	// type conversion of int literal (10) to myInt type
	m := myInt(10)
	// m: 10
	fmt.Println("m:", m)
	// isEven: true
	fmt.Println("isEven:", m.isEven())
	m.Double()
	// m: 20
	fmt.Println("m:", m)
}
