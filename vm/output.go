package vm

import (
	"io"
	"os"
)

var (
	StdIn    io.Reader
	StdOut   io.Writer
	StdErr   io.Writer
	StdDebug io.Writer
	StdBC    io.Writer
)

func init() {
	StdIn = os.Stdin
	StdOut = os.Stdout
	StdErr = os.Stderr
	StdDebug = os.Stdout
	StdBC = os.Stdout
}
