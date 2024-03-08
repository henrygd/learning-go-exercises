// Write a function called fileLen that has an input parameter of type string and
// returns an int and an error. The function takes in a filename and returns the
// number of bytes in the file. If there is an error reading the file, return the error.
// Use defer to make sure the file is closed properly.

package main

import (
	"fmt"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)

	defer func() {
		file.Close()
	}()

	if err != nil {
		return 0, err
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	size := stat.Size()

	return int(size), nil
}

func main() {
	size, err := fileLen("main.go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File size:", size)
}
