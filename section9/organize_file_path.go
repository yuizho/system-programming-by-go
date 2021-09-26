package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	// "${GOPATH}/src/hoge" -> "/home/user/src/hoge"
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	// to clean path
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	// to clean path
	fmt.Println(Clean2("~/.emacs.d/init.el"))
	fmt.Println(Clean2("${HOME}/config"))

	// to absolute path
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)

	// to relative path
	relpath, _ := filepath.Rel(
		"/usr/local/go/src",
		"/usr/local/go/src/path/filepath/path.go",
	)
	fmt.Println(relpath)
}
