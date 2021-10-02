package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func main() {
	// share by parameter
	go sub1(10)

	// share by a capture of closure
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
