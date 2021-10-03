package main

import "fmt"

func main() {
	s := make([]int, 1, 2)
	fmt.Println(&s[0], len(s), cap(s))

	s = append(s, 1)
	fmt.Println(&s[0], len(s), cap(s))

	s = append(s, 2)
	// head address is changed
	fmt.Println(&s[0], len(s), cap(s))
}
