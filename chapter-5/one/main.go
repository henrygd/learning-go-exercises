// The simple calculator program doesn’t handle one error case: division by zero.
// Change the function signature for the math operations to return both an int and
// an error. In the div function, if the divisor is 0, return errors.New("division
// by zero") for the error. In all other cases, return nil. Adjust the main function
// to check for this error.

package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	add = func(i, j int) (int, error) { return i + j, nil }
	sub = func(i, j int) (int, error) { return i - j, nil }
	mul = func(i, j int) (int, error) { return i * j, nil }
	div = func(i, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("division by zero")
		}
		return i / j, nil
	}
)

func main() {
	x, _ := add(2, 3)
	fmt.Println(x)

	x, _ = sub(3, 2)
	fmt.Println(x)

	x, _ = mul(2, 3)
	fmt.Println(x)

	x, err := div(3, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
}
