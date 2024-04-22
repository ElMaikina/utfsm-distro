package main

import (
	"context"
	"log"
	"net"
	pb "main/proto"
	"google.golang.org/grpc"
)

type myGemaEstrategicaServer struct {
	pb.UnimplementedGemaEstrategicaServer
}

func (s myGemaEstrategicaServer) SolicitarM(context.Context, *pb.Solicitud) (*pb.Respuesta, error) {
	log.Printf("Llego un mensaje nuevo!")
	return &pb.Respuesta{Rpta: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myGemaEstrategicaServer{}
	pb.RegisterGemaEstrategicaServer(serverRegister, service)
	
	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
