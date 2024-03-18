// Create a struct named Person with three fields: FirstName and LastName of
// type string and Age of type int. Write a function called MakePerson that
// takes in firstName, lastName, and age and returns a Person. Write a second
// function MakePersonPointer that takes in firstName, lastName, and age and
// returns a *Person. Call both from main. Compile your program with go build
// -gcflags="-m". This both compiles your code and prints out which values
// escape to the heap. Are you surprised about what escapes?

package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	// this is popped from stack bc we return the data itself
	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return person
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	// this is moved to heap bc we returned only a ref to it, not the data itself
	personPointer := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return &personPointer
}

func main() {
	MakePerson("Homer", "Simpson", 43)
	MakePersonPointer("Homer", "Simpson", 43)
}
