package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedNodesCommunicationServer
}

func (s *server) NotifyDecision(ctx context.Context, in *pb.Decision) (*pb.Response, error) {
	filename := in.ClientIp + "_" + in.Floor + ".txt"
	fmt.Println("Writing data")
	writeFile(filename, in.Option)
	return &pb.Response{Res: ""}, nil
}

func main() {
	fmt.Print("Starting server...\n")
	StartServer()
}

func StartServer() {
	listener, err := net.Listen("tcp", "200.1.22.237:50053")
	if err != nil {
		log.Fatalf("Error al crear servidor GRPC en la dirección %s: %s", "200.1.22.237:50053", err)
	}
	grpcServer := grpc.NewServer()
	service := &server{}
	pb.RegisterNodesCommunicationServer(grpcServer, service)
	log.Println("Server listening on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func writeFile(filename string, text string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
