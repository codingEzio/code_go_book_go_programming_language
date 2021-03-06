// A TCP server that periodically writes the time
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	fmt.Printf("Now you can test this by running `nc localhost %d`\n", *port)

	// listener, err := net.Listen("tcp", "localhost:8080")
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// If you don't add this `go` keyword, this program will be a
		// sequential one, that is, it can only accept one conn at a time.
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected
		}

		time.Sleep(1 * time.Second)
	}
}
