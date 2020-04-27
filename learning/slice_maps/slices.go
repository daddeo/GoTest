package main

import (
	"fmt"
)

func main() {
	test1()
	fmt.Println("-----------------------------------")
	test2()
	fmt.Println("-----------------------------------")
	test3()
}

func test1() {
	// slice declaration, start length of the slice (0), slice capacity (how much memory to prealloc for a slice)
	s := make([]string, 0)
	// length of s: 0
	fmt.Println("length of s:", len(s))
	s = append(s, "hello")
	// length of s: 1
	fmt.Println("length of s:", len(s))
	// contents of s[0]: hello
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	// contents of s[0]: goodbye
	fmt.Println("contents of s[0]:", s[0])

	// has two elements in the slice, s2[0] and s2[1]
	s2 := make([]string, 2)
	// contents of s2[0]:
	fmt.Println("contents of s2[0]:", s2[0])
	// adds to the length of the slice, in this case s2[2]
	s2 = append(s2, "hello")
	s2[1] = "goodbye"
	// contents of s2[1]:
	fmt.Println("contents of s2[1]:", s2[1])
	// contents of s2[2]: hello
	fmt.Println("contents of s2[2]:", s2[2])
	// length of s2: 3
	fmt.Println("length of s2:", len(s2))
}

func test2() {
	// mutable
	s3 := []string{"a", "b", "c"}

	// 0 a
	// 1 b
	// 2 c
	for k, v := range s3 {
		fmt.Println(k, v)
	}

	s4 := s3[0:2]
	// s4: [a b]
	fmt.Println("s4:", s4)
	s3[0] = "d"
	// s3: [d b c]
	fmt.Println("s3:", s3)
	// slices are reference types and behave like pointers
	// s4 is using s3 and since s3[0] changed s4 also represents that change
	// both point to the same memory, so a change to one is a change to the other
	// s4: [d b]
	fmt.Println("s4:", s4)

	// since no value was specified gets the 0 values and the zero value for a slice (like a pointer) is Nil
	var s5 []string
	s5 = s3
	s5[1] = "camel"
	// s3: [d camel c]
	fmt.Println("s3:", s3)

	modSlice(s3)
	// s3[0]: hello
	fmt.Println("s3[0]:", s3[0])
	// s3: [hello camel c]
	fmt.Println("s3:", s3)
}

func modSlice(s []string) {
	s[0] = "hello"
}

func test3() {
	uniHello := "👋 🌎"

	// is a copy of the uniHello string as bytes
	bytes := []byte(uniHello)
	// bytes: [240 159 145 139 240 159 140 142]
	fmt.Println("bytes:", bytes)

	// is a copy of the uniHello string as runes
	runes := []rune(uniHello)
	// runes: [128075 32 127758]
	fmt.Println("runes:", runes)
	runes[1] = 'a'
	// runes: [128075 97 127758]
	fmt.Println("runes:", runes)

	// as we were working with a copy of the original uniHello it has been unaltered
	// uniHello: 👋 🌎
	fmt.Println("uniHello:", uniHello)
}
