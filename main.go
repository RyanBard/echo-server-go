package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Print("Caller closed connection, returning...")
				return
			}
			log.Printf("Error reading string! err=%s", err.Error())
			return
		}
		_, err = fmt.Fprintln(conn, line)
		if err != nil {
			log.Printf("Error writing string! err=%s", err.Error())
			return
		}
	}
}

func main() {
	log.Print("Starting on port 8080...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080! err=%s", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection! err=%s", err.Error())
		}
		log.Print("Connection accepted, handling...")
		go handleConnection(conn)
	}
}
