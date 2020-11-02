package server

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func sendRequest(request, network string) string {
	conn, err := net.Dial("tcp", network)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't connect to:", network)
		return ""
	}

	conn.Write([]byte(request))
	response := make([]byte, maxBytes)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return ""
	}

	responseStr := string(response[:n])

	return responseStr
}

func TestServerFunctional(t *testing.T) {
	addr := net.JoinHostPort("127.0.0.1", "1456")

	testDict := map[string]string{
		"messi":  "the GOAT",
		"fun":    "enjoyment, amusement, or lighthearted pleasure",
		"secret": "something you tell everybody to tell nobody",
	}
	tests := []struct {
		request, response string
	}{
		{
			request:  "GET messi\n",
			response: "ANSWER the GOAT\n",
		},
		{
			request:  "GET fun\n",
			response: "ANSWER enjoyment, amusement, or lighthearted pleasure\n",
		},
		{
			request:  "ALL\n",
			response: "OK defined words: messi, fun, secret\n",
		},
		{
			request:  "SET fb foo bar\n",
			response: "OK definition for 'fb' was set\n",
		},
		{
			request:  "GET fb\n",
			response: "ANSWER foo bar\n",
		},
		{
			request:  "GET doesntexist\n",
			response: "ERROR 'doesntexist' definition not found\n",
		},
		{
			request:  "CLEAR\n",
			response: "OK all definitions cleared\n",
		},
		{
			request:  "ALL\n",
			response: "ERROR no definitions exist\n",
		},
	}

	server := NewServer(testDict)
	go server.Listen(addr)
	time.Sleep(time.Second)

	for _, tc := range tests {
		t.Run(tc.request, func(t *testing.T) {
			response := sendRequest(tc.request, addr)
			assert.Equal(t, tc.response, response)
		})
	}
}
