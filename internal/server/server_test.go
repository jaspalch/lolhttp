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
