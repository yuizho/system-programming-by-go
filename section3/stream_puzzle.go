package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	sytem       = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(sytem, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i1 := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(a, s, c, i1, i2)

	io.Copy(os.Stdout, stream)
}
