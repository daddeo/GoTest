package main

//import "fmt"
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Outside a function, every statement begins with a keyword (var, func, and so on) and so
// the := construct is not available.

// define three package-level boolean variables
var b1, b2, b3 bool

// define package-level int variables with initializers
var i, j int = 1, 2

func main() {
	bytesWritten, e := fmt.Printf("hello, world\n")
	fmt.Printf("bytes written: %d (includes newline)\n", bytesWritten)
	fmt.Println(e)

	bytesWritten, _ = fmt.Printf("hello, world\n")
	fmt.Printf("bytes written: %d (includes newline)\n", bytesWritten)

	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	name1 := "\n provides a new line."
	fmt.Println(name1)
	name2 := "\\n provides a new line."
	fmt.Println(name2)
	name3 := `\n provides a new line.` // treat as literal
	fmt.Println(name3)
	name4 := `\n provides a new
	line
		if we
			REALLY
				wanted one.` // treat as literal, even across multiple lines
	// string as immutable, so the following line won't work
	// name4[22] = "s"
	fmt.Println(name4)
	fmt.Println(len(name4))

	// will error compilation
	//fmt.Println(math.pi)
	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))
	fmt.Println(add2(88, 19))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
	first, second := split(56)
	fmt.Printf("first %d, second %d\n", first, second)

	// define function-level integer
	var i1 int
	fmt.Println(i1, b1, b2, b3)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	s := "Hello, World!"
	s1 := s[0:5]
	s2 := s[7:12]
	s3 := s[:5]
	s4 := s[7:]
	s5 := "👋 🌎"
	s6 := s5[:1]
	s7 := s5[2:]
	b1 := s[0]
	b2 := s[4]

	fmt.Println(s, b1, string(b1), b2, string(b2))
	fmt.Println(s, len(s), s3, len(s3), s4, len(s4))
	fmt.Println(s, s1, s2, s3, s4)
	// wrong cause it's looking at s5 as bytes, but s5 is UTF-8 and each unicode character can be 1-4 bytes each, hence len = 9
	fmt.Println(s5, len(s5), s6, len(s6), s7, len(s7))

	// rune represents a unicode character is same as int32 (4 bytes)
	s10 := "Hello, "
	var r rune = 127757
	s11 := s10 + string(r)
	r = '🌎'
	s12 := s10 + string(r)
	fmt.Println(s10, len(s10), s11, len(s11), s12, len(s12))

	short_var_decl()
}

// x int, y int can be shortened to x, y int
func add(x int, y int) int {
	return x + y
}

// the type comes after the variable name
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // known as a "naked" return
}

func short_var_decl() {
	var a, b int = 1, 2

	// Inside a function, the := short assignment statement can be used in place of a var declaration
	// with implicit type.
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(a, b, k, c, python, java)
}
