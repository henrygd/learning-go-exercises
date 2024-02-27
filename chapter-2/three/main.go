// Write a program with three variables, one named b of type byte, one named
// smallI of type int32, and one named bigI of type uint64. Assign each variable
// the maximum legal value for its type; then add 1 to each variable. Print out their
// values.

package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = math.MaxInt32
	bigI = math.MaxUint64

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
