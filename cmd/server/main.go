package main

import (
	"log"
	"ratelimiter/internal/server"
)

const Version = "0.0.1"

func main() {
	log.Printf("Starting rate limiter grpc server, v%s...\n", Version)
	server.Serve()
}
