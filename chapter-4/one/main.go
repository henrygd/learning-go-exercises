// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.

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

	fmt.Println(intSlice)
}
