package main

import (
	"io"
	"os"
	"strings"
)

func copyN(source io.Writer, dest io.Reader, length int64) (written int64, err error) {
	return io.Copy(
		source,
		io.LimitReader(dest, length),
	)
}

func main() {
	reader := strings.NewReader("1234567890")
	copyN(os.Stdout, reader, 5)
}
