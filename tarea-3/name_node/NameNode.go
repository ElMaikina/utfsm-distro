package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"net"
)

type server struct {
	pb.UnimplementedNodesCommunicationServer
}

func (s *server) NotifyDecision(ctx context.Context, in *pb.Decision) (*pb.Response, error) {
	fmt.Println("received data")
	SendDecisionToDataNode(in.ClientIp, in.Option, in.Floor)
	return &pb.Response{Res: ""}, nil
}

func main() {
	fmt.Print("Starting server...\n")
	StartServer()
}

func StartServer() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al crear servidor GRPC en la dirección %s: %s", ":50051", err)
	}
	grpcServer := grpc.NewServer()
	service := &server{}
	pb.RegisterNodesCommunicationServer(grpcServer, service)
	log.Println("Server listening on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func SendDecisionToDataNode(ip string, decision string, floor string) {

	conn, err := grpc.Dial(":50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewNodesCommunicationClient(conn)

	client.NotifyDecision(context.Background(), &pb.Decision{Option: decision, ClientIp: ip, Floor: floor})

	conn.Close()
}
