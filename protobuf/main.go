package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/k-stz/goscratch/protobuf/person"
	"google.golang.org/protobuf/proto"
)

// core "k8s.io/api/core/v1"
//pb "./person.proto"

// type Person struct {
// 	name string
// 	age  int
// }

func main() {
	bob := &person.Person{
		Name: "Protobob",
		Age:  200,
	}
	fmt.Println("Person:", bob)

	b, err := proto.Marshal(bob)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
 
	fo, err := os.Create("person-protobuf.pb")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	// Write protobuf bytes to file
	if _, err := w.Write(b); err != nil {
		panic(err)
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}
