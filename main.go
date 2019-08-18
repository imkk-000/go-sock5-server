package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const PORT = 2000

func newServer(port int) net.Listener {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func handleConn(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	var buffer []byte
	var err error
	for err != io.EOF {
		buffer, _, err = reader.ReadLine()
		log.Printf("Recv: %s", string(buffer))
	}
	log.Printf("Client %s disconnected", c.RemoteAddr().String())
}

func startListener(port int) {
	// create listener for server
	server := newServer(port)
	defer server.Close()
	log.Printf("start listener on port %d", port)

	for {
		// receive connection from client
		client, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Has client connected from %s", client.RemoteAddr().String())

		// create new routine for handle client
		go handleConn(client)
	}
}

func main() {
	startListener(PORT)
}
