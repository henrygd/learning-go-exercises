// Write a program that defines a variable named greetings of type slice of
// strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは",
// and "Привіт". Create a subslice containing the first two values; a second subslice
// with the second, third, and fourth values; and a third subslice with the fourth and
// fifth values. Print out all four slices.

package main

import "fmt"

func main() {
	var greeting = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subsliceOne := greeting[:2]
	subsliceTwo := greeting[1:4]
	subsliceThree := greeting[3:]
	fmt.Println(subsliceOne)
	fmt.Println(subsliceTwo)
	fmt.Println(subsliceThree)
}
