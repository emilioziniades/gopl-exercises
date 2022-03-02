package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan<- string // an outgoing message channel
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

const timeout = 1000 * time.Second

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			welcome := "current chat participants: "
			for cli := range clients {
				welcome += fmt.Sprintf("%s, ", cli.name)
			}
			cli.ch <- welcome
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	ch <- "enter name:"
	input.Scan()
	who := input.Text()

	cli := client{ch, who}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		fmt.Fprintln(conn, "Timeout: 10 seconds of inactivity. Closing connection")
		conn.Close()
	}()

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
