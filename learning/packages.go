package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// just a test example, as ROT13 is one of the worst encryption types available
// just rotates every character by 13 letters, but can be used to encode and decode
func rot13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}
	return in
}

func main() {
	s := "This is a test 123 😃"
	s2 := strings.Map(rot13, s)
	fmt.Println(s2)
	s3 := strings.Map(rot13, s2)
	fmt.Println(s3)

	// strings does not contain functions for working with unicode characters, need to use
	// unicode/utf8
	usingRunes()
	getTime()
}

func usingRunes() {
	s := "👋🌎"
	fmt.Println(s)
	fmt.Println(len(s))
	// RuneCountInString tells us how many unicode characters are in the string, not bytes
	fmt.Println(utf8.RuneCountInString(s))
}

func getTime() {
	t := time.Now() // type time.Time
	fmt.Println(t)
	nanos := t.UnixNano() // type int64
	fmt.Println(nanos)
}
