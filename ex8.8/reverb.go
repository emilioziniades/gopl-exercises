// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var in = make(chan string)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleInput(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		in <- input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func handleConn(c net.Conn) {
	for {
		select {
		case i := <-in:
			go echo(c, i, 1*time.Second)
		case <-time.After(5 * time.Second):
			fmt.Fprintln(c, "5 second timeout, closing connection")
			c.Close()
			return
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleInput(conn)
		go handleConn(conn)
	}
}
