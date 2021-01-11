package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jaswraith/lolhttp/internal/server"
)

func handleErr(e error) {
	if e == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %v\n", e)
	os.Exit(1)
}

func main() {
	dict := map[string]string{
		"messi":  "the GOAT",
		"fun":    "enjoyment, amusement, or lighthearted pleasure",
		"secret": "something you tell everybody to tell nobody",
	}

	address := "127.0.0.1"
	port := "1456"

	srv := server.NewServer(dict)

	err := srv.Listen(net.JoinHostPort(address, port))
	handleErr(err)
}
