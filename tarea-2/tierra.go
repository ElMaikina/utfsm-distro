package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

type myGemaEstrategicaServer struct {
	mustEmbedUnimplementedGemaEstrategicaServer
}

func (s myGemaEstrategicaServer) SolicitarM(context.Context, *Solicitud) (*Respuesta, error) {
	return Respuesta{ rpta: 6 }, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myGemaEstrategicaServer{}
	RegisterGemaEstrategicaServer(serverRegister, service)
	
	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	
}
