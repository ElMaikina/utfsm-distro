package main

import (
	"log"
	pb "main/proto"
)

type myGemaEstrategicaServer struct {
	pb.UnimplementedGemaEstrategicaServer
}

func (s myGemaEstrategicaServer) SolicitarM(context.Context, *pb.Solicitud) (*pb.Respuesta, error) {
	log.Printf("Llego un mensaje nuevo!")
	return &pb.Respuesta{Rpta: 1}, nil
}

func main() {

}
