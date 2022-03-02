package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var a = flag.Int("a", 256, "choose which SHA algorithm: 256, 384 or 512")
	flag.Parse()

	if *a != 256 && *a != 384 && *a != 512 {
		log.Fatalf("invalid SHA algorithm %d. choose either SHA 256, 384 or 512", a)
	}

	var buf bytes.Buffer
	buf.ReadFrom(os.Stdin)
	in := bytes.TrimSuffix(buf.Bytes(), []byte("\n"))
	var out []byte

	switch *a {
	case 256:
		hash := sha256.Sum256(in)
		out = hash[:]
	case 384:
		hash := sha512.Sum384(in)
		out = hash[:]
	case 512:
		hash := sha512.Sum512(in)
		out = hash[:]

	}
	fmt.Printf("%X\n", out)
}
