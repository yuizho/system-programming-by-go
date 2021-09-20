package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start!!!")
	timer := time.After(time.Second * 10)
	<-timer
	fmt.Println("Done!")
}
