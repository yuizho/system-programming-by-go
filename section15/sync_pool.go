package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	// when this pool is blank, this function will be returned
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	pool.Put("manualy added: 3")
	// pool values are cleared
	runtime.GC()
	fmt.Println(pool.Get())
}
