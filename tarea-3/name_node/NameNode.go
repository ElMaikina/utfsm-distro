package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"math/rand"
	"net"
)

// TODO: asignar ip correctas
var ipDataNode_1 = "localhost:50053"
var ipDataNode_2 = "localhost:50054"
var ipDataNode_3 = "localhost:50055"

type server struct {
	pb.UnimplementedNodesCommunicationServer
}

func (s *server) NotifyDecision(ctx context.Context, in *pb.Decision) (*pb.Response, error) {
	fmt.Println("received data")
	node := rand.Intn(3)
	if node == 0 {
		SendDecisionToDataNode(in.ClientIp, in.Option, in.Floor, ipDataNode_1)
	} else if node == 1 {
		SendDecisionToDataNode(in.ClientIp, in.Option, in.Floor, ipDataNode_2)
	} else {
		SendDecisionToDataNode(in.ClientIp, in.Option, in.Floor, ipDataNode_3)
	}
	return &pb.Response{Res: ""}, nil
}

func main() {
	fmt.Print("Starting server...\n")
	StartServer()
}

func StartServer() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al crear servidor GRPC en la dirección %s: %s", ":50052", err)
	}
	grpcServer := grpc.NewServer()
	service := &server{}
	pb.RegisterNodesCommunicationServer(grpcServer, service)
	log.Println("Server listening on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func SendDecisionToDataNode(ip string, decision string, floor string, ipDataNode string) {

	conn, err := grpc.Dial(ipDataNode, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewNodesCommunicationClient(conn)
	client.NotifyDecision(context.Background(), &pb.Decision{Option: decision, ClientIp: ip, Floor: floor})
	conn.Close()
}
