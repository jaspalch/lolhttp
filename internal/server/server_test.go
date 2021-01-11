package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	reqMethods := []string{"GET", "SET", "CLEAR", "ALL"}
	tests := []map[string]string{
		{},
		{
			"compiler": "converts source code into machine language",
		},
		{
			"messi":  "the GOAT",
			"fun":    "enjoyment, amusement, or lighthearted pleasure",
			"secret": "something you tell everybody to tell nobody",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test dictionary: %v", tc), func(t *testing.T) {
			s := NewServer(tc)
			assert.Equal(t, reqMethods, s.allowedMethods)
			assert.Equal(t, tc, s.dict)
			for _, m := range reqMethods {
				_, ok := s.methodHandlers[m]
				assert.True(t, ok, m+" methodHandler not set")
			}
		})
	}
}

func TestListen(t *testing.T) {
	tests := []struct {
		s    *Server
		addr string
		err  error
	}{
		{
			s: &Server{
				allowedMethods: []string{"GET"},
			},
			addr: "127.0.0.1",
			err:  fmt.Errorf("No handler function set for method 'GET'"),
		},
		{
			s:    NewServer(map[string]string{}),
			addr: "bad address",
			err:  fmt.Errorf("listen tcp: address bad address: missing port in address"),
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test error: %v", tc.err), func(t *testing.T) {
			err := tc.s.Listen(tc.addr)
			assert.Equal(t, tc.err.Error(), err.Error())
		})
	}
}

func TestRegisterHandler(t *testing.T) {
	tests := []struct {
		method string
		err    error
	}{
		{
			method: "BAD",
			err:    fmt.Errorf("Method 'BAD' is not a valid method"),
		},
		{
			method: "GET",
			err:    nil,
		},
	}

	srv := NewServer(map[string]string{})

	for _, tc := range tests {
		t.Run(fmt.Sprintf("method to register: %v", tc.method), func(t *testing.T) {
			err := srv.registerHandler(tc.method, func(s *Server, args []string) string { return "" })
			assert.Equal(t, tc.err, err)
		})
	}
}
