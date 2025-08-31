package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Environment configs
	port := os.Getenv("GOMEMDB_PORT")
	if port == "" {
		port = ":6379"
	}

	// Create new server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error to start tcp server (%v): %v", port, err)
	}
	log.Printf("tcp listening on port: %v", port)

	// Accepting connections
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("error to accept new connections: %v", err)
	}
	defer conn.Close()

	for {

		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))

	}

}
