package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	path := os.Getenv("PATH")
	fmt.Printf("env variable PATH: %s\n", path)

	for _, path := range filepath.SplitList(path) {
		excecpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(excecpath)
		if !os.IsNotExist(err) {
			fmt.Println(excecpath)
			return
		}
	}
	os.Exit(1)
}
