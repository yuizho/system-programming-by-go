package main

import (
	"fmt"
	"io"
	"os"

	colorable "github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)

func main() {
	var stdOut io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		stdOut = colorable.NewColorableStdout()
	} else {
		stdOut = colorable.NewNonColorable(os.Stdout)
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Fprintln(stdOut, "stdin: terminal")
	} else {
		fmt.Fprintln(stdOut, "stdin: pipe")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Fprintln(stdOut, "stdout: terminal")
	} else {
		fmt.Fprintln(stdOut, "stdout: pipe")
	}
	if isatty.IsTerminal(os.Stderr.Fd()) {
		fmt.Fprintln(stdOut, "stderr: terminal")
	} else {
		fmt.Fprintln(stdOut, "stderr: pipe")
	}
}
