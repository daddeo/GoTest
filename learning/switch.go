package main

import (
	"fmt"
	"os"
)

func main() {
	var word string
	if len(os.Args) > 1 {
		word = os.Args[1]
	} else {
		word = "hello"
	}

	test_if_else(word)
	test_switch(word)
	test_switch2(word)
}

func test_if_else(word string) {
	if word == "hello" {
		fmt.Println("Hi yourself.")
	} else if word == "goodbye" {
		fmt.Println("So long.")
	} else if word == "greetings" {
		fmt.Println("Saluations.")
	} else {
		fmt.Println("I don't know what you said")
	}
}

func test_switch(word string) {
	greeting := "greetings"
	goodbye := "goodbye"
	switch l := len(word); word {
	case "hi":
		fmt.Println("informal.")
		fallthrough // allows case statement to fall through to next case as well
	case "hello":
		fmt.Println("Hi yourself.")
	case "farewell": // legal, just is a no-op match
	case goodbye, "bye":
		fmt.Println("So long.")
	case greeting:
		fmt.Println("Saluations.")
	default:
		fmt.Println("I don't know what you said, but it was", l, "characters long")
	}
}

func test_switch2(word string) {
	c := "crackerjack"

	switch l := len(word); {
	case word == "hi":
		fmt.Println("informal.")
		fallthrough // allows case statement to fall through to next case as well
	case word == "hello":
		fmt.Println("Hi yourself.")
	case l == 1:
		fmt.Println("I don't know any one letter words.")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either", c, "or it is 2-9 characters long.")
	default:
		fmt.Println("I don't know what you said, but it was", l, "characters long")
	}
}
