package main

import (
	"fmt"
	"io"
	"os"
)

type person struct {
	age  int
	name string
}

func modifyFails(p person) {
	p.name = "Doug Dimmadome"
	fmt.Printf("%p\n", p)
	fmt.Println(p)

}

// fileLen takes a filename and returns the number of bytes in the file
func fileLen(filename string) (count int, err error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	buffer := make([]byte, 2048)
	for {
		c, err := fd.Read(buffer)
		count += c
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
	}

}

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + s
	}
}

func main() {
	fn := prefixer("Dimma")
	s := fn("Doug")
	fmt.Println(fn(fn("Dome")), s)
}
