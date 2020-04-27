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
// go run learning/goerrors/errors3.go
// output:
// e is nil: true
// in main, e is nil: true
// me is nil: true
// in main, me is nil: false
//
// me (interface) is only nil if the interface has no underlying type assigned to it
// in notReallyNil(), me is nil since not underlying type was assigned
// BUT
// when it is returned and assigned to me, me now has a underlying non-nil value
//
// NOTE: THIS APPLIES TO ALL INTERFACE TYPES!!!
// very counter-intuitive and not cool!
// IMPORTANT: If defining your own error types, need to be sure that you never define a variable
// to be your own error type otherwise it will no return nil EVEN IF there is no error occurred.
//

package main

import (
	"fmt"
)

// ErrorEx2 an error struct
type ErrorEx2 struct {
	A       int
	B       int
	Message string
}

func (me *ErrorEx2) Error() string {
	return fmt.Sprintf("Values: %d and %d produced error %s", me.A, me.B, me.Message)
}

func reallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)
	return e
}

func notReallyNil() error {
	var me *ErrorEx2
	fmt.Println("me is nil:", me == nil)
	return me
}

func main() {
	e := reallyNil()
	fmt.Println("in main, e is nil:", e == nil)

	me := notReallyNil()
	fmt.Println("in main, me is nil:", me == nil)
}
