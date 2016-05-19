package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	//var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	//debug 为 false 时 out 的 动态类型 不为 nil 而是 *bytes.Buffer,所以 out 不为 nil
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
