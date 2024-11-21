package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func handleClient(clientFd int) {
	defer syscall.Close(clientFd)

	reader := bufio.NewReader(syscallConn(clientFd))

	request, err := parseHTTPRequest(reader)
	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return
	}

	log.Printf("Received %s request for %s", request.Method, request.Path)

	response := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        "Content-Length: 13\r\n" +
        "Connection: close\r\n\r\n" +
        "Hello, World!"

	if _, err := syscall.Write(clientFd, []byte(response)); err != nil {
		log.Printf("Error writing to client: %v", err)
	}

	log.Printf("Sent: %v", response)

	time.Sleep(1 * time.Second)
}

func syscallConn(fd int) *os.File {
	return os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
}

func createSocket() (int, error) {
	fd, error := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)

	if error != nil {
		return -1, error
	}

	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		return -1, err
	}

	return fd, nil
}

func acceptConnections(fd int) {
	for {
		clientFd, _, err := syscall.Accept(fd)
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		peerName, err := syscall.Getpeername(clientFd)
		if err != nil {
			log.Printf("Failed to get peer name: %v", err)
			continue
		}
		log.Printf("Accepted connection from %v", peerName)
		
		
		
		handleClient(clientFd)
	}
}

