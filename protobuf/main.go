package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/k-stz/goscratch/protobuf/myconfigmap"
	"google.golang.org/protobuf/proto"
	core "k8s.io/api/core/v1"
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
	fooDataMap := make(map[string]string)
	fooDataMap["name"] = "My awesome ConfigMap protobuf test"
	fooCM := &myconfigmap.MyConfigMap{
		Data: fooDataMap,
	}
	fmt.Println("myconfigmap", fooCM)

	b, err := proto.Marshal(fooCM)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	writeProtobufToFile(b, "my-protobuf.pb")

	readProtobufBytes := readFileToByteSlice("my-protobuf.pb")
	// lifeform.Person protobuf can be unmarshalled into lifeform.Animal!
	barCM := new(myconfigmap.MyConfigMapSameEnumAs1)
	proto.Unmarshal(readProtobufBytes, barCM)
	fmt.Println("MyConfigMapSameEnumAs1: ", barCM)

	quxCMdifferntEnum := new(myconfigmap.MyConfigMap2)
	proto.Unmarshal(readProtobufBytes, quxCMdifferntEnum)
	fmt.Println("MyConfigMap2:           ", quxCMdifferntEnum)

	bothConfigMap := new(myconfigmap.DataBinaryDataConfigMap)
	proto.Unmarshal(readProtobufBytes, bothConfigMap)
	fmt.Println("DataBinaryDataConfigMap:", bothConfigMap)

	// COnfigMap
	// Data map[string]string `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
	coreCMdataMap := make(map[string]string)
	cm := core.ConfigMap{
		Data: coreCMdataMap,
	}
	fmt.Println("core CM:", cm)
}
