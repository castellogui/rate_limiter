package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func Serve() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	log.Printf("addr %s\n", os.Getenv("GRPC_SERVER"))
	listener, err := net.Listen("tcp", os.Getenv("GRPC_SERVER"))
	if err != nil {
		panic(fmt.Sprintf("error while trying to start tcp listner: %s", err))
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("error while trying to serve grpc server: %s", err))
	}
}
