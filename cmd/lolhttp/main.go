package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func handleErr(e error) {
	if e == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %v\n", e)
	os.Exit(1)
}

func handleConnection(conn net.Conn) {
	interval := 3 * time.Second

	for t := range time.Tick(interval) {
		_, err := conn.Write([]byte(fmt.Sprintf("Hello whoever is connected!\nThe time is: %v\n", t)))
		handleErr(err)
	}

	conn.Close()
}

func main() {
	address := "127.0.0.1"
	port := "1456"

	fmt.Println("Starting server on", net.JoinHostPort(address, port))
	listener, err := net.Listen("tcp", net.JoinHostPort(address, port))
	handleErr(err)

	for {
		conn, err := listener.Accept()
		handleErr(err)

		fmt.Println("Accepted connection!")
		go handleConnection(conn)
	}
}
