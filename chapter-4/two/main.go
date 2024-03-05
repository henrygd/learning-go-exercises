// Loop over the slice you created in exercise 1. For each value in the slice, apply the
// following rules:
// a. If the value is divisible by 2, print “Two!”
// b. If the value is divisible by 3, print “Three!”
// c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// d. Otherwise, print “Never mind”.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := make([]int, 100)

	for i := range intSlice {
		intSlice[i] = rand.Intn(100)
	}

	for _, v := range intSlice {
		switch {
		case v == 0:
			fmt.Println(v, "Never mind")
		case v%2 == 0:
			fmt.Println(v, "Two!")
		case v%3 == 0:
			fmt.Println(v, "Three!")
		case v%6 == 0:
			fmt.Println(v, "Six!")
		default:
			fmt.Println(v, "Never mind")
		}
	}
}
