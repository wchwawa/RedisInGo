package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 128)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			return
		}
		request, _ := parseRequest(buf)
		response, _ := parseResponseCommand(request)
		_, err = conn.Write(response)

		if err != nil {
			log.Printf("Failed to write to connection: %v", err)
			return
		}
	}
}
