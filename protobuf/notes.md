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

# Unmarshalling
Very important: When unmarshalling a protobuf into a struct, the structs tag is crucial for it decides where which data is put into the struct.
For example when Creating the protobuf based on a struct that with the following field:
```go
	Data map[string]string `protobuf:"bytes,2,rep,name=Data,proto3" 
```
Where the "2" after `"bytes,2,...` indicates the enum of the Message 
in the `*.proto` file.

Now when you use a different Struct for unmarshalling, that has multiple
fields, make sure that the field has the same enum in its struct tag. 
For example:
```go
	Data map[string]string `protobuf:"bytes,3,...`
```
Will _not_ unmarshal the data into the Data map becaue the enums don't match. 
But instead you can create a protoc-struct that uses multiple fields, as long as the field you care about has the current enum matching the marshalled protobuf data. So the following would unmarshal it successfully, even though the struct has more fields then the marshalled protobuf:
```go
type DataBinaryDataConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       map[string]string `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BinaryData map[string]string `protobuf:"bytes,3,rep,name=BinaryData,proto3" json:"BinaryData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
```

Here the `Data` field Data containing a map would properly receive the data!