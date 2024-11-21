package main

import (
	"log"
	"syscall"
)

const port = 8080

func main() {
	fd, err := createSocket()

	if err != nil {
		log.Fatalf("Error creating socket: %v", err)
	}

	defer syscall.Close(fd)
	
	addr := syscall.SockaddrInet4{Port: port}
	copy(addr.Addr[:], []byte{127, 0, 0 , 1})

	if err := syscall.Bind(fd, &addr); err != nil {
		log.Fatalf("Error binding socket: %v", err)
	}

	log.Printf("Server listening on port %d", port)

	if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatalf("Error listening on socket: %v", err)
	}
    

	acceptConnections(fd)
}