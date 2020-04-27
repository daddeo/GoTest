// Caesar Cipher

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var offset uint8 = 5

func main() {
	name := "Caesar Cipher"
	fmt.Printf("name: %s\n", name)
	fmt.Printf("character type: %T\n", name[0])

	fmt.Printf("character codes: ")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%v,", name[i])
	}
	fmt.Printf("\n")

	var encodedStr strings.Builder
	encodedBytes := bytes.Buffer{}
	for i := 0; i < len(name); i++ {
		encodedBytes.WriteByte(encode(name[i]))
		encodedStr.WriteByte(encode(name[i]))
	}
	// buf.WriteString("first string")
	// buf.WriteString("second string")
	encodedText := encodedBytes.String()
	fmt.Printf("encodedText: %s\n", encodedText)
	fmt.Printf("encodedStr: %s\n", encodedStr.String())

	decodedBytes := bytes.Buffer{}
	for i := 0; i < len(encodedText); i++ {
		decodedBytes.WriteByte(decode(encodedText[i]))
	}
	fmt.Printf("decodedText: %s\n", decodedBytes.String())

	// encodedString.ReadAt()
	var decodedStr strings.Builder
	for i := 0; i < encodedStr.Len(); i++ {
		decodedStr.WriteByte(decode(encodedStr.String()[i]))
	}
	fmt.Printf("decodedStr: %s\n", decodedStr.String())

	// var encodedString strings.Reader
	var finalStr strings.Builder
	r := strings.NewReader(encodedStr.String())
	size := 8 // could be len(name) for whole buffer
	buffer := make([]byte, size)
	for {
		n, err := r.Read(buffer)
		fmt.Printf("n: {%v}, err: {%v}\n", n, err)
		fmt.Printf("buffer: {%v}\nbuffer[:n]: {%q}\n", buffer, buffer[:n])
		if err == io.EOF {
			break
		}
		finalStr.Write(buffer)
	}
	fmt.Printf("finalStr: {%s}\n", finalStr.String())

}

func encode(character uint8) uint8 {
	return character + offset
}

func decode(character uint8) uint8 {
	return character - offset
}
