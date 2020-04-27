package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100) // 100 >= seed >= 0
	b := rand.Intn(100)
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, "is bigger")
}
