package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("initialize")
}

func main() {
	var once sync.Once
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
