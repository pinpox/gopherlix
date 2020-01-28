package main

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

// Define basic settings of the server. This will be exported to a config file
// at some point.

type GopherServer struct {
	Port    string
	Domain  string
	Host    string
	RootDir string
	run     bool
	signals chan bool
}

func NewGopherServer(port, domain, host, root string) GopherServer {
	return GopherServer{Port: port, Domain: domain, Host: host, RootDir: root, run: false, signals: make(chan bool)}

}

func (server *GopherServer) Run() {
	server.run = true

	// Listen for incoming connections.
	l, err := net.Listen("tcp", server.Host+":"+server.Port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}

	// Close the listener when the application closes.

	defer l.Close()

	log.Infof("Listening on %s:%s", server.Host, server.Port)
	for {

		select {
		case stop := <-server.signals:
			if stop {
				log.Info("Stop signal received, stopping Server")
				break
			}
		default:
		}

		// Listen for an incoming connection.
		conn, err := l.Accept()
		log.Println("Accepted connection from:", conn.RemoteAddr())
		if err != nil {
			log.Warn("Error accepting: ", err.Error())
		}

		// Handle connections in a new goroutine.
		go server.handleRequest(conn)
		if err != nil {
			log.Println(err)
		}
	}
}

func (server *GopherServer) parseRequest(req string, reqLen int) (string, error) {
	reqPath := strings.TrimSuffix(req, "\r\n")
	log.Info("Got request: " + reqPath)

	savePath, err := server.getSavePath(reqPath)

	if err != nil {
		log.Fatal(err)
	}

	listing, err := server.createListing(savePath)

	if err != nil {
		log.Fatal(err)
	}

	return listing, nil
}

// Handles incoming requests.
func (server *GopherServer) handleRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return err
	}

	log.Println("Read ", reqLen, " bytes from ", conn.RemoteAddr())
	log.Println(string(buf[:reqLen]))

	response, err := server.parseRequest(string(buf[:reqLen]), reqLen)

	if err != nil {
		log.Println("Error parsing request:", err.Error())
		return err
	}

	// Send a response back to person contacting us.
	log.Println("Sending response to: ", conn.RemoteAddr())
	log.Println(response)
	conn.Write([]byte(response))

	// Close the connection when you're done with it.
	log.Info("Closing conntion to: ", conn.RemoteAddr())
	conn.Close()
	return nil
}
