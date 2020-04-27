// there are no exceptions in Go
// Go allows for multiple return values from a single method or function
// by convention, the last return value from a method or function indicates whether there was an error
// the last parameter is of type:
//
// type error interface {
//    Error() string
// }
//

// run from command line with:
// go run learning/goerrors/errors1.go 10 2
// get:
// 10 / 2 == 5, 10 % 2 == 0
//
// go run learning/goerrors/errors1.go asdf 2
// output:
// Invalid argument #1
// error: strconv.ParseInt: parsing "asdf": invalid syntax
// exit status 1
//
// go run learning/goerrors/errors1.go 10 asdf
// output:
// Invalid argument #2
// error: strconv.ParseInt: parsing "asdf": invalid syntax
// exit status 1
//
// go run learning/goerrors/errors1.go 10 0
// output:
// There was an error: Cannot divide by zero
// exit status 1

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	return a / b, a % b, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Expected two input parameters")
		os.Exit(1)
	}
	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid argument #1\nerror:", err)
		os.Exit(1)
	}
	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Invalid argument #2\nerror:", err)
		os.Exit(1)
	}
	// have to cast because ParseInt technically returns int64 values
	div, mod, err := divAndMod(int(a), int(b))
	if err != nil {
		fmt.Println("There was an error:", err)
		os.Exit(1)
	}

	fmt.Printf("%d / %d == %d, %d %% %d == %d\n", a, b, div, a, b, mod)
}
