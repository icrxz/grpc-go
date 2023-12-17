package config

import (
	"log"
	"net"
)

func InitListener() net.Listener {
	addr := "localhost:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen address %s. Error: %v\n", addr, err)
	}

	log.Printf("Started listening %s...\n", addr)

	return listener
}
