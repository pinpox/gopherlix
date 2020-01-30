package main

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GopherServer holds the basic information of the server. This includes
// connection parameters and the server root directory
type GopherServer struct {

	// Port to listen on, normally 70
	Port string

	// Domain to which the requests will be made to. This will be used e.g.
	// in links
	Domain string

	// Host to bind to, most likely localhost or a specific IP
	Host string

	ServerRoot *GopherServerRoot

	// Control server main loop. Setting this to false or sending a signal
	// to the channel will result in stopping the server
	run     bool
	signals chan bool
}

// NewGopherServer is used to create a new server. It returns a server, that is not running yet
func NewGopherServer(port, domain, host, root string) GopherServer {

	rootDir, err := NewGopherServerRoot(root)

	// We can't continue without a working server root
	if err != nil {
		log.Fatal(err)
	}

	return GopherServer{
		Port:       port,
		Domain:     domain,
		Host:       host,
		ServerRoot: rootDir,
		run:        false,
		signals:    make(chan bool),
	}
}

// Run starts the server. It will listen for connections until the stop signal
// is send via the signals channel.
func (server *GopherServer) Run() {
	server.run = true

	// Listen for incoming connections.
	l, err := net.Listen("tcp", server.Host+":"+server.Port)

	if err != nil {
		// If any error occures here, print it and quit. We can't continue at
		// this point
		log.Fatal("Error listening:", err.Error())
	}

	// Close the listener when the application closes.
	defer l.Close()

	log.Infof("Listening on %s:%s", server.Host, server.Port)

	// Main loop, this will run until we receive the stop signal or an error
	// occurs
	for {

		// Read from the signals channel in a non-blocking fashion. In case we
		// get a signal, stop the server printing out an informational message
		select {
		case stop := <-server.signals:
			if stop {
				log.Info("Stop signal received, stopping Server")
				break
			}
		default:
		}

		// Listen for an incoming connection. This will block until a connection is made.
		conn, err := l.Accept()

		// Log accepted connection with ip address.
		log.Println("Accepted connection from:", conn.RemoteAddr())
		if err != nil {
			log.Warn("Error accepting: ", err.Error())
		}

		// Handle connections in a new goroutine. If any errors occur during
		// handling of the requests, don't quit but close the connection and
		// continue listening
		go server.handleRequest(conn)
		if err != nil {
			log.Error(err)
		}
	}
}

// parseRequest parses the request and decides what the reponse should be. It
// is then returned as a simple string to be send by handleRequest
func (server *GopherServer) parseRequest(req string) (string, error) {

	// Log request
	log.Info("Request: \""+replaceCRLF(req), "\"")

	// Trim trailing \r\n characters
	reqPath := strings.TrimSuffix(req, "\r\n")

	// Request path is allowed, try to create a response
	listing, err := server.createListing(reqPath)

	// If the file was not found or any other error occured, return an error
	if err != nil {
		log.Warn("Error creating response for: \"", reqPath, "\"")
		return "", err
	}

	// Everything is fine, return the response
	return listing, nil
}

// Handles incoming requests.
func (server *GopherServer) handleRequest(conn net.Conn) error {

	// Make sure we close the connection after using it
	defer conn.Close()

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return err
	}

	// Create a response from the request
	response, err := server.parseRequest(string(buf[:reqLen]))

	// If the request could not be parsed or any error occured, just send an
	// error message and return an error
	if err != nil {
		conn.Write([]byte("Invalid request"))
		return err
	}

	// Send response
	_, err = conn.Write([]byte(response))

	// Return an error, if any occured while writing to the connection. Should
	// be nil in most cases
	return err
}
