package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%s, %d, %f", "os.File example", 1, 1.3)
	file.Close()
}
