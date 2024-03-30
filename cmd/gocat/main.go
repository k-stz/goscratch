package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

// It is a common pattern in Go, for a function that allocates resources,
// like `open()`, to also provide a closure that cleans up the resource.
// The `closer` closure, "closes over" the allocated filedescriptor.
// The caller can then defer the closure ensuring the file is closed upon use
// or an error reported if any
// 
// Other langs use blocks like try/catch/finally (java/JavaScript) 
// or begin/rescure/ensure (Ruby) for this purpose instead.
func open(filename string) (fd *os.File, closer func(), err error) {
	fd, err = os.Open(filename)
	// We provide a closure to the caller, that can be defer-ed
	// it will ensure the file is closed and report errors if any
	closer = func() {
		fmt.Println("closing file!")
		err = fd.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	return fd, closer, err
}

func main() {
	slog.Info("Arguments passed", "os.Args", os.Args, "len", len(os.Args))
	var filename string
	if len(os.Args) < 2 {
		filename = "test.txt"
		slog.Warn("Not enough args passed. Usage: gocat <file>")
		slog.Warn("Using default filename.", "filename", filename)
	} else {
		filename = os.Args[1]
	}
	// Path to file is relative to current working dir (os.Getwd())
	cwd, _ := os.Getwd()
	slog.Info("os.Getwd()", "cwd", cwd)
	fd, closer, err := open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	buffer := make([]byte, 2048)
	i := 0
	for {
		i++
		fmt.Println("iteration num:", i)
		count, err := fd.Read(buffer)
		fmt.Println("err after read:", err)
		s := buffer[0:count]
		os.Stdout.Write(s)
		if err != nil {
			// EOF error gets returned when _no_ bytes were read.
			// So when a file has any content, you won't get
			// an EOF on the first read!
			if err != io.EOF {
				log.Fatal(err)
			}
			break // EOF error
		}
	}

}
