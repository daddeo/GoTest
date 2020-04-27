package main

/*
BASIC TYPES
---------------------------------------------------------
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32, represents a Unicode code point
	float32 float64
	complex64 complex128

usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
*/

import (
	"fmt"
	"math"
	"math/cmplx"
)

// a variable block like imports
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type {%T} = {%v}\n", ToBe, ToBe)
	fmt.Printf("type {%T} = {%v}\n", MaxInt, MaxInt)
	fmt.Printf("type {%T} = {%v}\n", z, z)

	//var hello = "Hello", world = "World!"
	var hello, world = "Hello", "World!"
	fmt.Printf("%s %s\n", hello, world)
	fmt.Println(hello, world) // a space is automatically added between vars for println

	zeroValues()
	typeConvert()
	typeInference()
	constants()
	numericConstant()
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("zero values ==> int = %v, float64 = %v, bool = %v, string = %q\n", i, f, b, s)
}

func typeConvert() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// could be also defined as:
	//x, y := 3, 4
	//f := math.Sqrt(float64(x*x + y*y))
	//z := uint(f)

	fmt.Printf("x: %v, y: %v, sqrt(x^2 + y^2): %f, z: %v\n", x, y, f, z)
}

func typeInference() {
	var x string
	y := x // y is an string
	fmt.Printf("y is of type %T\n", y)

	b := true         // boolean
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

const Pi = 3.14

// Constants cannot be declared using the := syntax.
func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// Numeric constants are high-precision values.
// An untyped constant takes the type needed by its context.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstant() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// overflows int
	// fmt.Println(needInt(Big))
	// overflows uint64
	// large_number := uint64(Big)
}
