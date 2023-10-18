package structs

import "fmt"

/*
	Structs
	Can be used to store a collection of fields.
	Similar to classes.
*/

type Person struct {
	name string
	age  int
}

func CreatePerson(name string, age int) Person {
	person := Person{name, age}

	fmt.Printf("Created person with name %v and age %v\n", person.name, person.age)
	return person
}
