package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("fileInfo")
	fmt.Printf("    filename: %v\n", info.Name())
	fmt.Printf("    filesize: %v\n", info.Size())
	fmt.Printf("    updated: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("    is directory?: %v\n", info.Mode().IsDir())
}
