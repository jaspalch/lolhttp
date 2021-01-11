package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		s     string
		found bool
	}{
		{
			s:     "not found",
			found: false,
		},
		{
			s:     "found",
			found: true,
		},
	}

	list := []string{"found"}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("test string: %v", tc.s), func(t *testing.T) {
			found := find(tc.s, list)
			assert.Equal(t, tc.found, found)
		})
	}
}
