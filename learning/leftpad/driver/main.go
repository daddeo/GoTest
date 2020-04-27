package main

import (
	"fmt"
	"test/learning/leftpad"
	"test/learning/leftpad/cycle"
)

/*
	to run from command line use:
	> go run leftpad\driver\main.go
*/
func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("¿Cómo está?", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
	fmt.Println(leftpad.FormatRune("¿Cómo está?", 15, '😀'))
	fmt.Println(cycle.FormatDouble("15.2221", 15))
}
