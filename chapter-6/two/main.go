// Write two functions. The UpdateSlice function takes in a []string and a
// string. It sets the last position in the passed-in slice to the passed-in string. At
// the end of UpdateSlice, print the slice after making the change. The GrowSlice
// function also takes in a []string and a string. It appends the string onto the
// slice. At the end of GrowSlice, print the slice after making the change. Call these
// functions from main. Print out the slice before each function is called and after
// each function is called. Do you understand why some changes are visible in main
// and why some changes are not?

package main

import "fmt"

func UpdateSlice(strSlice []string, str string) {
	// will update bc capacity is unchanged
	length := len(strSlice)
	strSlice[length-1] = str
	fmt.Println(strSlice)
}

func GrowSlice(strSlice []string, str string) {
	// added but capacity for main variable is unchanged so doesn't show
	strSlice = append(strSlice, str)
	fmt.Println(strSlice)
}

func main() {
	strSlice := []string{"Michigan", "Washington"}
	str := "Alabama"
	fmt.Println(strSlice)
	UpdateSlice(strSlice, str)
	fmt.Println(strSlice)
	GrowSlice(strSlice, str)
	fmt.Println(strSlice)
}
