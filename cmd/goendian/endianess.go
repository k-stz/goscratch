package main

import (
	"fmt"
	"log/slog"
	"unsafe"
)

func main() {
	// new(type) returns pointer to zero-value instance of given type
	// x := new(int)
	// fmt.Println(x)
	// fmt.Println("*x:", *x)

	// Is Big endian?
	fmt.Println("Big Endian Test")
	i := int32(10)
	up := unsafe.Pointer(&i)
	slog.Info("Vars defined", "i", i, "&i", &i, "unsafe pointer to i:", up)
	b0 := (*byte)(up)
	b1 := (*byte)(unsafe.Add(up, 1))
	b2 := (*byte)(unsafe.Add(up, 2))
	b3 := (*byte)(unsafe.Add(up, 3))
	fmt.Println("byte 0:", b0, "*byte 0:", *b0)
	fmt.Println("byte 1:", b1, "*byte 1:", *b1)
	fmt.Println("byte 2:", b2, "*byte 2:", *b2)
	fmt.Println("byte 3:", b3, "*byte 3:", *b3)
	fmt.Print("Thus, Machine uses ")
	if *b0 == 10 && *b3 == 0 {
		fmt.Println("Little Endian!")
	} else if *b0 == 0 && *b3 == 10 {
		fmt.Println("Big Endian!")
	} else {
		fmt.Println("neiter big or little endian... Somethings wrong!")
	}

	//ptr = &i
	//unsafe.Add(&i, 1)

}
