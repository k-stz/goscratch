package main

import "fmt"

type MyFunc func(i int)
type MyFunc2 func(i, j int)

type printer interface {
	printNum(i int)
}

func (m MyFunc) printNum(i int) {
	fmt.Println("Calling printNum with MyFunc: m(i)")
	m(i)
}

func (m MyFunc2) printNum(i int) {
	fmt.Println("Calling printNum with MyFunc: m(i)")
	m(i, i)
}

func myHandler(pr printer, stuff string) {
	fmt.Println("myHandler called", pr, stuff)
	pr.printNum(100)
}

func main() {
	var mf MyFunc = func(i int) {
		fmt.Println("I'm mf then: ", i)
	}
	var mf2 MyFunc2 = func(i, j int) {
		fmt.Println("I'm MyFunc2! ", i, j)
	}
	mf.printNum(22)
	mf2.printNum(33)
	myHandler(mf, "blub")
}
