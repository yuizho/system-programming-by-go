package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		go func(name string) {
			mutex.Lock()
			defer mutex.Unlock()
			cond.Wait()
			fmt.Println(name)
		}(name)
	}
	fmt.Println("ready")
	time.Sleep(time.Second)
	fmt.Println("start!")
	cond.Broadcast()
	time.Sleep(time.Second)
}
