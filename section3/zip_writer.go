package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("this is read by strigs.NewReader")
	file, _ := os.Create("file.zip")
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, _ := zipWriter.Create("newfile.txt")
	io.Copy(writer, reader)
}
