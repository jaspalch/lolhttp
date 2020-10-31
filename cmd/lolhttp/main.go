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

func main() {
	address := "127.0.0.1"
	port := "1456"

	fmt.Println("Starting server on", net.JoinHostPort(address, port))
	listener, err := net.Listen("tcp", net.JoinHostPort(address, port))
	handleErr(err)

	for {
		_, err := listener.Accept()
		handleErr(err)

		fmt.Println("Accepted connection!")
	}
}
