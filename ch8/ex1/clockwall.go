package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 12, 2, ' ', 0)
	args := parseArgs(os.Args[1:])
	clocks := connectClocks(args)
	printLocations(tw, clocks)

	for {
		for _, c := range clocks {
			fmt.Fprintf(tw, "%v\t", c.Time())
		}
		fmt.Fprint(tw, "\r")
		tw.Flush()
	}

}

type Clock struct {
	location string
	host     string
	conn     net.Conn
	s        *bufio.Scanner
}

func (c *Clock) Time() string {
	c.s.Scan()
	if err := c.s.Err(); err != nil {
		return fmt.Sprint(err)
	}
	return c.s.Text()
}

func connectClocks(args map[string]string) []Clock {
	clocks := make([]Clock, len(args))
	i := 0
	for place, host := range args {
		conn, err := net.Dial("tcp", host)
		defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(conn)
		clocks[i] = Clock{place, host, conn, s}
		i++
	}
	return clocks
}

func printLocations(tw *tabwriter.Writer, clocks []Clock) {
	fmt.Println("")
	for _, c := range clocks {
		fmt.Fprintf(tw, "%v\t", c.location)
	}
	fmt.Fprint(tw, "\n")
	for _, _ = range clocks {
		fmt.Fprint(tw, "--------\t")
	}
	fmt.Fprint(tw, "\n")
}

func parseArgs(args []string) (res map[string]string) {
	res = make(map[string]string)
	for _, arg := range args {
		a := strings.Split(arg, "=")
		res[a[0]] = a[1]
	}
	return
}
