package server

import (
	"fmt"
	"sort"
	"strings"
)

func getHandler(s *Server, args []string) string {
	if len(args) != 1 {
		return "ERROR incorrect number of arguments for GET request (want 1)"
	}

	word := args[0]
	def, ok := s.dict[word]
	if !ok {
		return fmt.Sprintf("ERROR '%v' definition not found", word)
	}

	return fmt.Sprintf("ANSWER %v", def)
}

func setHandler(s *Server, args []string) string {
	if len(args) < 2 {
		return "ERROR incorrect number of arguments for SET request (want >= 2)"
	}

	word := args[0]
	s.dict[word] = fmt.Sprintf(strings.Join(args[1:], " "))

	return fmt.Sprintf("OK definition for '%v' was set", word)
}

func clearHandler(s *Server, args []string) string {
	if len(args) != 0 {
		return "ERROR incorrect number of arguments for CLEAR request (want 0)"
	}

	s.dict = make(map[string]string)
	return "OK all definitions cleared"
}

func allHandler(s *Server, args []string) string {
	if len(args) != 0 {
		return "ERROR incorrect number of arguments for ALL request (want 0)"
	}

	if len(s.dict) == 0 {
		return "ERROR no definitions exist"
	}

	var words []string
	for k := range s.dict {
		words = append(words, k)
	}

	sort.Strings(words)
	return "OK defined words: " + strings.Join(words, ", ")
}
