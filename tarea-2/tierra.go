package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GemaEstrategicaServer struct{}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	serverRegister := grpc.NewServer()

}
