package main

import "fmt"

type Adder struct {
	i int
}

type Arithmer struct {
	// Adder is an "embedded field"!
	// Fields will be directly accessible
	// and Adder's methods will be promoted to Arithmer
	// and thus an Airthmer struct can use the AddTo()-method!
	Adder
	j int
}

func (a Adder) AddTo(value int) int {
	fmt.Println("Adder's AddTo() called value:", value)
	return a.i + value
}

func (a Adder) Double(value int) int {
	// Here we call AddTo on a, an instance of Adder
	// Such that the Adder's AddTo() will be called!
	return a.AddTo(value * 2)
}

func (a Arithmer) Double(value int) int {
	// Here we call AddTo on a, an instance of Adder
	// SUch that the Adder's AddTo() will be called!
	return a.AddTo(value * 2)
}

func (a Arithmer) AddTo(value int) int {
	fmt.Println("Arithmer's AddTo() called value:", value)
	return a.i + value
}

func main() {
	//test()
	a := Adder{1}
	// mv is a method value
	mv := a.AddTo
	var myAdder func(int) int = a.AddTo
	// this is a method expression: creating a function from the type itself
	me := Adder.AddTo
	fmt.Println(a, a.AddTo(4))
	fmt.Println("method values:", mv(5), myAdder(8), me(Adder{1}, 13))
	arith := Arithmer{Adder{3}, 4}
	fmt.Println("## Embedding test:")
	//fmt.Println("arith:", arith, arith.AddTo(55))
	fmt.Println("Double:", arith.Double(10))

}
