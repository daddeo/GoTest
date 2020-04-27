package main

import "fmt"

func main() {
	var vals [3]int // once declared cannot be redeclared, eg has length 3 now, cannot be a different length later
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	// length of the array is considered part of the type and therefore cannot assign [3]int to a [4]int
	// var vals2 [4]int = vals
	// slices are used instead of arrays cause of this limitation

	fmt.Println(vals, vals[0], vals[1], vals[2])
}
