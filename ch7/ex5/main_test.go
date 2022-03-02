package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitedReader(t *testing.T) {
	buf := &bytes.Buffer{}
	data := "Hello To You"
	reader := LimitReader(strings.NewReader(data), 5)
	n, err := buf.ReadFrom(reader)
	if read := buf.String(); err != nil || n != 5 || read != "Hello" {
		t.Fatalf("n=%d err=%s read=%s", n, err, read)
	}
}
