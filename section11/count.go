package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
		fmt.Fprintln(os.Stderr, "stderr")
	}
}
