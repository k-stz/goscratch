package main

import (
	"fmt"
)

type Greeter interface {
	Hi()
}

type GreeterLeaver interface {
	Greeter
	Bye()
}

type Num struct {
	i int
}

func (n Num) Hi() {
	fmt.Println("hi, im a Num:", n.i)
}

func (n Num) Bye() {
	fmt.Println("Bye, Num needs to leave:", n.i)
}

func Salute(g Greeter) {
	g.Hi()
	fmt.Println("and I Salute you!")
}

func Adios(gl GreeterLeaver) {
	gl.Bye()
	fmt.Println("I'm leaving")
}

func main() {
	num := Num{22}
	num.Hi()
	Salute(num)
	Adios(num)

	str := fmt.Sprintf("%d %s\n", 1, "hi")
	fmt.Println("string is: ", str)

	type MyInt int
	var a any
	var mine MyInt = 44
	a = mine
	switch j := a.(type) {
	case nil: 
		fmt.Println("Type is nil")
	case int: 
		fmt.Println("Type is int")
	case MyInt: 
		fmt.Println("Type is MyInt")
	case string: 
		fmt.Println("Type is string")
	default: 
		fmt.Println("No idea what i is, so j is of type any", j)
	}

}