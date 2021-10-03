package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// init
	fmt.Println(time.Now().String())
	fmt.Println(time.Date(2020, time.October, 2, 22, 46, 0, 0, time.Local))
	fmt.Println(time.Parse(time.Kitchen, "11:30AM"))

	// calc
	fmt.Println(time.Now().Add(3 * time.Hour))
	fmt.Println(time.Now().Add(3 * -time.Hour))

	fileInfo, _ := os.Stat("/etc/sysctl.d")
	fmt.Printf("the file was created before %v\n", time.Now().Sub(fileInfo.ModTime()))

	fmt.Println(time.Now().Round(time.Hour))
}
