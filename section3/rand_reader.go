package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	buffer := make([]byte, 1024)
	rand.Read(buffer)
	fmt.Println(len(buffer))
	fmt.Println(buffer)
}
