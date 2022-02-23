package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		io.WriteString(c, "> ")
		r, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		cmd := strings.TrimRight(r, "\n")
		switch {
		case cmd == "pwd":
			wd, err := os.Getwd()
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Fprintln(c, wd)
		case cmd == "ls":
			wd, err := os.Getwd()
			if err != nil {
				log.Print(err)
				return
			}
			files, err := ioutil.ReadDir(wd)
			if err != nil {
				log.Print(err)
				return
			}
			for _, f := range files {
				if f.IsDir() {
					fmt.Fprint(c, "(D)")
				}
				fmt.Fprintln(c, f.Name())
			}
		case strings.HasPrefix(cmd, "cd"):
			cmds := strings.Split(cmd, " ")
			if len(cmds) < 2 {
				os.Chdir("/Users/emilioziniades")
			} else {
				os.Chdir(cmds[1])
			}
		case strings.HasPrefix(cmd, "get"):
			cmds := strings.Split(cmd, " ")
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Print(err)
				file.Close()
			}
			io.Copy(c, file)
			file.Close()
		case cmd == "close":
			return
		default:
			fmt.Fprintln(c, "err: unknown command")
		}
	}
}
