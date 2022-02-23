package main

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	s := "Hello"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Logf("n=%d err=%s", n, err)
		t.Fail()
	}

	if b.String() != s {
		t.Logf("%s != %s", b.String(), s)
	}

}

func TestNewReaderWithHTML(t *testing.T) {
	s := "<h1> Hello there <a> Emilio </a> </h1>"
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		t.Logf("err=%s", err)
		t.Fail()
	}
	fmt.Println(doc)
}
