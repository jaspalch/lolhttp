package main

import (
	"fmt"
	"net"
	"os"
)

func handleErr(e error) {
	if e == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %v\n", e)
	os.Exit(1)
}

func handleConnection(conn net.Conn) {
	numMsgs := 5

	// Send welcome message
	_, err := conn.Write([]byte("Hello whoever is connected!\n"))
	handleErr(err)

	for i := 0; i < numMsgs; i++ {
		// Receive reply from client
		reply := make([]byte, 4096)
		_, err = conn.Read(reply)
		handleErr(err)

		// Send back reply to client
		_, err = conn.Write(append([]byte("You said:\n"), reply...))
		handleErr(err)
	}

	_, err = conn.Write([]byte("Terminating connection...\n"))
	handleErr(err)
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
