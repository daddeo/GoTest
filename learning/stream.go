package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	println("-----------------------------------------------------")
	testRand()
	println("-----------------------------------------------------")
	testJoin()
	println("-----------------------------------------------------")
	printFloat()
	sprintFloat()
	fmt.Println(RandString(10))
	fmt.Println(RandString(30))
	fmt.Println(strings.HasPrefix("something", "some"))
	fmt.Println(strings.HasSuffix("something", "thing"))
	fmt.Println(strings.Contains("something", "me"))

	var str string = "this is a string"
	toAndFromSlice(str)

}

func testRand() {
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(randString())
	}

	fmt.Println(b.String())
}

func testJoin() {
	var strs []string

	for i := 0; i < 1000; i++ {
		strs = append(strs, randString())
	}

	fmt.Println(strings.Join(strs, ""))
}

func randString() string {
	// Pretend to return a random string
	return "abc-123-"
}

// https://golang.org/pkg/strconv/#pkg-index
func printFloat() {
	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
}

func sprintFloat() {
	i := 123
	t := fmt.Sprintf("We are currently processing ticket number %d.", i)
	fmt.Println(t)
}

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

// https://golang.org/ref/spec#Conversions
func toAndFromSlice(s string) {
	fmt.Printf("s: {%s}\n", s)
	var b []byte
	b = []byte(s)
	fmt.Printf("b: {")
	for i := range b {
		fmt.Printf("%v", string(b[i]))
	}
	fmt.Printf("}\n")
	s = string(b)
	fmt.Printf("s: {%s}\n", s)
}
