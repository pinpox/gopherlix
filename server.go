package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Define basic settings of the server. This will be exported to a config file
// at some point.
const (
	CONN_HOST   = "localhost"
	CONN_DOMAIN = "localhost"
	CONN_PORT   = "8000"
	CONN_TYPE   = "tcp"
	SERVER_ROOT = "./content"
)

func main() {

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {

		// Listen for an incoming connection.
		conn, err := l.Accept()
		fmt.Println("Accepted connection from:", conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
		if err != nil {
			log.Println(err)
		}
	}
}

func parseRequest(req string, reqLen int) (string, error) {
	reqPath := strings.TrimSuffix(req, "\r\n")
	fmt.Println("Creating listing for \"" + reqPath + "\"")

	listing, err := createListing(reqPath)

	if err != nil {
		return "", nil
	}

	return listing, nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Read ", reqLen, " bytes from ", conn.RemoteAddr())
	fmt.Println(string(buf[:reqLen]))

	response, err := parseRequest(string(buf[:reqLen]), reqLen)

	// Send a response back to person contacting us.
	fmt.Println("Sending response to: ", conn.RemoteAddr())
	fmt.Println(response)
	conn.Write([]byte(response))

	// Close the connection when you're done with it.
	fmt.Println("Closing conntion to: ", conn.RemoteAddr())
	conn.Close()
	return nil
}
