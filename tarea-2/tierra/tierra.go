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

type Tierra struct {
	AT        int // Cantidad de munición AT en inventario
	MP        int // Cantidad de munición MP en inventario
	MaxAT     int // Máximo de munición AT que se puede almacenar
	MaxMP     int // Máximo de munición MP que se puede almacenar
	grpcAddr  string
}

func (t *Tierra) GenerarMunicion() {
	for {
		time.Sleep(5 * time.Second)
		t.AT += 10 // Genera 10 AT cada 5 segundos
		t.MP += 5  // Genera 5 MP cada 5 segundos
		// Verifica si el inventario supera el límite máximo y ajusta si es necesario
		if t.AT > t.MaxAT {
			t.AT = t.MaxAT
		}
		if t.MP > t.MaxMP {
			t.MP = t.MaxMP
		}
	}
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

	Tierra tierra;
	GenerarMunicion()
}
