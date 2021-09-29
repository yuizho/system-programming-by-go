package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	count := exec.Command("./count")
	stdout, _ := count.StdoutPipe()
	stderr, _ := count.StderrPipe()
	var m sync.Mutex

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m.Lock()
			fmt.Printf("(stdout) %s\n", scanner.Text())
			m.Unlock()
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			m.Lock()
			fmt.Printf("(stderror) %s\n", scanner.Text())
			m.Unlock()
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}
}
