package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var portt = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	fmt.Printf("Now you can test this by running `nc localhost %d`\n", *portt)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portt))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn)
	}
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConn2(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1500*time.Millisecond) // 1.5s
	}
	conn.Close()
}
