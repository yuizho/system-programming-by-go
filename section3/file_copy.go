package main

import (
	"io"
	"os"
)

func main() {
	reader, err := os.Open("file_copy.go")
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	new_file_name := "new_file_copy.txt"
	if len(os.Args) > 1 {
		new_file_name = os.Args[1]
	}

	writer, err := os.Create(new_file_name)
	if err != nil {
		panic(err)
	}
	defer writer.Close()
	io.Copy(writer, reader)
}
