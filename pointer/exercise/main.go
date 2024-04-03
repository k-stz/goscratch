package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName, lastName, age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	person := MakePerson(firstName, lastName, age)
	return &person
}

func main() {
	// Compile prgoram with -gcflags="-m" to get result of escape analysis 
	// for which values will escape to the heap
	steve := MakePerson("Steve", "Stacker", 55)
	// ezekiel's pointed to value will escape the stack, because after
	// the stack frame pops, it needs to be referred to. So the escape analysis
	// places it on the heap.
	ezekiel := MakePersonPointer("Ezekiel", "Escapist", 44)
	// Steve Stacker will also escape the stack! That's solely due to this Println,
	// as it expects an `..any`. The current Go compiler moves any argument that is
	// of an interface type to the heap.
	fmt.Println(steve)
	fmt.Println(ezekiel)
}
