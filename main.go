package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				logrus.Info("Caller closed connection, returning...")
				return
			}
			logrus.Warn("Error reading string! err=", err.Error())
			return
		}
		_, err = fmt.Fprintln(conn, line)
		if err != nil {
			logrus.Warn("Error writing string! err=", err.Error())
			return
		}
	}
}

func main() {
	logrus.Info("Starting on port 8080...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal("Failed to listen on port 8080! err=", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Warn("Failed to accept connection! err=", err.Error())
		}
		logrus.Info("Connection accepted, handling...")
		go handleConnection(conn)
	}
}
