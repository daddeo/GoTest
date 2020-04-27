package main

import (
	"fmt"
	"test/language/mapper"
)

func main() {
	fmt.Println(mapper.Greets("Howdy, what's new?"))
	fmt.Println(mapper.Greets("Comment allez vous?"))
	fmt.Println(mapper.Greets("Wie geht e Ihnen?"))
}
