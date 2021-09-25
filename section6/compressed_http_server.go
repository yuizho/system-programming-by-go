package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGzipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","),
		"gzip",
	) != -1
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	for {
		// configure timeout
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error) // down cast to get timeout error data
			if ok && neterr.Timeout() {   // timeout
				fmt.Println("Timeout")
				break
			} else if err == io.EOF { // socket was closed
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// Configures HTTP1.1 and ContentLength
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}
		if isGzipAcceptable(request) {
			content := "Hello World (gzipped)\n"
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = io.NopCloser(&buffer) // add Close method
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encodding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = ioutil.NopCloser(
				strings.NewReader(content),
			)
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip Server is runnning at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
