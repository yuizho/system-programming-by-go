package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	r, err := ioutil.ReadAll(teeReader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
	fmt.Println(buffer.String())
}
