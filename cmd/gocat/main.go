package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

func open(filename string) (fd *os.File, closer func(), err error) {
	fd, err = os.Open(filename)
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
