package main

import (
	"fmt"
)

func main() {
	skip := 3
	sum := 0
	for i := 0; i < 10; i++ {
		if i > 8 {
			break
		}
		if i == skip {
			continue
		}
		sum += i
	}
	fmt.Println(sum)

	i := 0
	for i < 10 {
		sum += i
		i += 1
	}
	fmt.Println(sum)

	i = 0
	for {
		sum += i
		i += 1
		if i < 10 {
			break
		}
	}
	fmt.Println(sum)

	s := "Hello World!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}

	s = "👋 🌎"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}

	a := 10
	if a == 1 {
		fmt.Println("on the first number")
	} else if a > 5 {
		if b := a / 2; b > 5 {
			fmt.Println("b is greater than 5")
		} else {
			fmt.Println("b is less than or equal to 5")
		}
		fmt.Println("a is bigger than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}
}
