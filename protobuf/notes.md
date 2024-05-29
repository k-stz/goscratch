# Define message formats in a `.proto` file

# How to use the protocol buffer compiler

# Use the Go protocol buffer API to write and read messages

# Problem domain of protobufs
- serialize go structs and share data with apps written for other platforms
- encode data into a single string (usually requires writing one-off encoding and parsing code)
- serialize data to XML. But XML is notoriously space itensive and encoding/decoding is a performance penalty on applications.
    - also navigating XML DOM tree is conserably more complplicated than navigating simple fields in a class normally would

=> protobufs are a flexible, efficient, automated solution to solve exactly this problem

# usage
1. you write a `.proto` description of the data structure you wish to store
2. from that, protobuf compiler creates a class that implements automatic encoding and parsing of the protobuf data with an efficient binary format
    - class provides getters/setters for the fields
    - takes care of reading/writing the protobuf as a unit
    - supports extension of the defined format over time in such a way that code can still read data encoded with the old format

# Example
Create a person.proto file:
```protobuf
syntax = "proto3";

option go_package = "./person";

message Person {
  string Name = 1;
  int64 Age = 2;
}
```
Thanks to the `option go_package` the `protoc` compiler will generate go file in that is in the go-package "person" for the message Person struct. And the whole API around it (functions on the struct)


Then compile it with protoc
```bash
# For example when the current dir contains a person.proto file:
protoc --go_out=. person.proto
```
The result will be a file `./perso/person.pb.go` containing the protobuf struct implmenenting code!
The struct herein finally implements the necessary interface to unmarshal/marshal a "Person" between the go struct "Person" and the protobuf wireformat for it.
