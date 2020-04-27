// Structs are not Objects
// there is no inheritance

package main

import (
	"encoding/json"
	"fmt"
)

// can be of any primitive type, other structs, functions, maps, slices, pointers
// structs that are capitalized are visible outside of their package like variables
// members within a struct starting with a capitalized letter are also visible outside of their package
// Foo is visible outside package
// foo is not visible outside package

// Foo is a test struct, visible outside of package wnd containing A (visible) and b (not).
type Foo struct {
	A int    // visible outside of package
	b string // not visible outside of package
}

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	fmt.Println("-----------------------------------")
	test3()
	fmt.Println("-----------------------------------")
	test4()
}

func test1() {
	// initialize struct to zero value which means all values inside are set to their zero value
	f := Foo{}
	// {0 }
	fmt.Println(f)

	// not common cause all fields must be initialized and non-descript
	g := Foo{10, "Hello"}
	// {10 Hello}
	fmt.Println(g)

	// more popular cause it is more descriptive, b is set to zero value
	h := Foo{
		b: "Goodbye",
	}
	// {0 Goodbye}
	fmt.Println(h)

	h.A = 1000
	// {1000 Goodbye}
	fmt.Println(h)
}

func test2() {
	// structs are value types and when copied are assigned to an entirely new memory location
	f := Foo{
		A: 20,
	}
	var f2 Foo // initialized to zero value of Foo
	f2 = f     // assigned to Foo
	f2.A = 100
	// f2.A: 100
	fmt.Println("f2.A:", f2.A)
	// f.A: 20
	fmt.Println("f.A:", f.A)

	// a pointer to struct foo (defined by f)
	var f3 *Foo = &f
	f3.A = 200
	// f3.A: 200
	fmt.Println("f3.A:", f3.A)
	// f.A: 200
	fmt.Println("f.A:", f.A)
}

// Bar is a test struct
type Bar struct {
	C string
	F Foo
}

// Baz is a test struct
type Baz struct {
	D   string
	Foo // embedded struct, not inheritance, allows for deligation -- contains a Foo as a field
}

func test3() {
	f := Foo{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: f}
	// f: {10 Hello}
	fmt.Println("f:", f)
	// b1: {Fred {10 Hello}}
	fmt.Println("b1:", b1)
	// b1.F.A: 10
	fmt.Println("b1.F.A:", b1.F.A)

	b2 := Baz{D: "Nancy", Foo: f}
	// b2: {Nancy {10 Hello}}
	fmt.Println("b2:", b2)
	// b2.A: 10
	fmt.Println("b2.A:", b2.A)

	var f2 Foo = b2.Foo
	// f2: {10 Hello}
	fmt.Println("f2:", f2)
}

// the Person struct contains a string literal called a struct tag

// Person is a test struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test4() {
	// a json snippet assigned to bob
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	// convert bob into a slice of bytes
	json.Unmarshal([]byte(bob), &b)
	// b: {Bob 30}
	fmt.Println("b:", b)
	// get a slice of bytes ingoring any errors
	bob2, _ := json.Marshal(b)
	// bob: {"name":"Bob","age":30}
	fmt.Println("bob:", string(bob2))
}
