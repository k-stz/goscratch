package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/k-stz/goscratch/protobuf/lifeform"
	"google.golang.org/protobuf/proto"
)

// core "k8s.io/api/core/v1"
//pb "./person.proto"

// type Person struct {
// 	name string
// 	age  int
// }

func writeProtobufToFile(b []byte, filename string) {
	fo, err := os.Create(filename)
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

func readFileToByteSlice(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	bob := &lifeform.Person{
		Name: "Bob the Human",
	}
	fmt.Println("bob:", bob)

	b, err := proto.Marshal(bob)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	writeProtobufToFile(b, "person-protobuf.pb")

	readProtobufBytes := readFileToByteSlice("person-protobuf.pb")
	// lifeform.Person protobuf can be unmarshalled into lifeform.Animal!
	someAnimal := new(lifeform.Animal)
	proto.Unmarshal(readProtobufBytes, someAnimal)
	fmt.Println("someAnimal:", someAnimal)
}
