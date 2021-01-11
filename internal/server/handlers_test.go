package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	srv := NewServer(map[string]string{
		"messi": "the GOAT",
	})

	tests := []struct {
		args []string
		resp string
	}{
		{
			args: []string{"nonexistent"},
			resp: "ERROR 'nonexistent' definition not found",
		},
		{
			args: []string{},
			resp: "ERROR incorrect number of arguments for GET request (want 1)",
		},
		{
			args: []string{"messi"},
			resp: "ANSWER the GOAT",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test request: %v", tc.args), func(t *testing.T) {
			resp := getHandler(srv, tc.args)
			assert.Equal(t, tc.resp, resp)
		})
	}
}

func TestSetHandler(t *testing.T) {
	srv := NewServer(map[string]string{})

	tests := []struct {
		args []string
		resp string
	}{
		{
			args: []string{},
			resp: "ERROR incorrect number of arguments for SET request (want >= 2)",
		},
		{
			args: []string{"messi", "the", "GOAT"},
			resp: "OK definition for 'messi' was set",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test request: %v", tc.args), func(t *testing.T) {
			resp := setHandler(srv, tc.args)
			assert.Equal(t, tc.resp, resp)
		})
	}
}

func TestClearHandler(t *testing.T) {
	srv := NewServer(map[string]string{})

	tests := []struct {
		args []string
		resp string
	}{
		{
			args: []string{"bad_arg"},
			resp: "ERROR incorrect number of arguments for CLEAR request (want 0)",
		},
		{
			args: []string{},
			resp: "OK all definitions cleared",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test request: %v", tc.args), func(t *testing.T) {
			resp := clearHandler(srv, tc.args)
			assert.Equal(t, tc.resp, resp)
		})
	}
}

func TestAllHandler(t *testing.T) {
	tests := []struct {
		srv  *Server
		args []string
		resp string
	}{
		{
			srv:  NewServer(map[string]string{}),
			args: []string{"bad_arg"},
			resp: "ERROR incorrect number of arguments for ALL request (want 0)",
		},
		{
			srv: NewServer(map[string]string{
				"messi":  "the GOAT",
				"fun":    "enjoyment, amusement, or lighthearted pleasure",
				"secret": "something you tell everybody to tell nobody",
			}),
			args: []string{},
			resp: "OK defined words: fun, messi, secret",
		},
		{
			srv:  NewServer(map[string]string{}),
			args: []string{},
			resp: "ERROR no definitions exist",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test request: %v", tc.args), func(t *testing.T) {
			resp := allHandler(tc.srv, tc.args)
			assert.Equal(t, tc.resp, resp)
		})
	}
}
