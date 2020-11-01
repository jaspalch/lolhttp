package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const maxBytes = 4096

type methodHandler func(*Server, []string) string

// Server defines the main dictionary server struct
type Server struct {
	listener       net.Listener
	dict           map[string]string
	allowedMethods []string
	methodHandlers map[string]methodHandler
}

// NewServer creates a new instance of Server and
// initializes the Server's word dictionary
func NewServer(dict map[string]string) *Server {
	server := Server{}

	server.dict = make(map[string]string)
	for k, v := range dict {
		server.dict[k] = v
	}
	server.allowedMethods = []string{"GET", "SET", "CLEAR", "ALL"}
	server.methodHandlers = make(map[string]methodHandler)

	server.registerHandler("GET", getHandler)
	server.registerHandler("SET", setHandler)
	server.registerHandler("CLEAR", clearHandler)
	server.registerHandler("ALL", allHandler)

	return &server
}

// Listen starts the Server on the provided address
func (s *Server) Listen(addr string) error {
	// Check if all method handlers are set
	for _, m := range s.allowedMethods {
		if _, ok := s.methodHandlers[m]; !ok {
			return fmt.Errorf("No handler function set for method '%v'", m)
		}
	}

	// Start listener on addr
	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	fmt.Println("Server is started and listening on", addr)

	// Accept and handle incoming connections
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when accepting connection: %v\n", err)
		}

		go s.mainHandler(conn)
	}
}

// mainHandler is the entrypoint for all incoming requests
// to the server. It checks requests' validity and passes
// them to the appropriate handler
func (s *Server) mainHandler(conn net.Conn) {
	defer conn.Close()

	// Read request
	request := make([]byte, maxBytes)
	numBytes, err := conn.Read(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when reading request: %v\n", err)
		return
	}

	// Check if request ends in a newline
	requestStr := string(request[:numBytes])
	if requestStr[len(requestStr)-1] != '\n' {
		errStr := "ERROR request is not newline terminated\n"
		conn.Write([]byte(errStr))
		fmt.Fprintf(os.Stderr, errStr)
		return
	}
	requestStr = strings.TrimSpace(requestStr)

	// Check if request method is valid
	requestArgs := strings.Split(requestStr, " ")
	method := requestArgs[0]
	if !find(method, s.allowedMethods) {
		errStr := fmt.Sprintf("ERROR method '%v' is not a valid method\n", method)
		conn.Write([]byte(errStr))
		fmt.Fprintf(os.Stderr, errStr)
		return
	}

	// Pass request to handler
	handlerFunc := s.methodHandlers[method]
	response := handlerFunc(s, requestArgs[1:])
	if err != nil {
		errStr := fmt.Sprintf("ERROR '%v'\n", err)
		conn.Write([]byte(errStr))
		fmt.Fprintf(os.Stderr, errStr)
		return
	}

	conn.Write([]byte(response + "\n"))
}

// Register method handlers to server
func (s *Server) registerHandler(method string, handler methodHandler) error {
	if !find(method, s.allowedMethods) {
		return fmt.Errorf("Method '%v' is not a valid method", method)
	}

	s.methodHandlers[method] = handler
	return nil
}
