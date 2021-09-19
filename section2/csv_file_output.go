package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test..csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(io.MultiWriter(file, os.Stdout))
	writer.WriteAll(
		[][]string{
			{"id", "name", "note"},
			{"1", "yui", "hoge"},
		},
	)
	writer.Flush()
}
