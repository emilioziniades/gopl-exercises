package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer
	newbuf, count := CountingWriter(&buf)
	data := []byte("Hello World")
	newbuf.Write(data)
	fmt.Println(*count)
}

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	new := &byteCounter{w, 0}
	return new, &new.written
}
