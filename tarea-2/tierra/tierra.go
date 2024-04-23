package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"time"
	"sync"
	pb "main/proto"
	"google.golang.org/grpc"
)

var maxAT int32 = 50
var maxMP int32 = 20
var AT int32 = 0
var MP int32 = 0

type myGemaEstrategicaServer struct {
	pb.UnimplementedGemaEstrategicaServer
}

func (s myGemaEstrategicaServer) SolicitarM(con context.Context, sol *pb.Solicitud) (*pb.Respuesta, error) {
	if (sol.AT > AT) {
		fmt.Printf("Recepcion de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA -- AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", sol.ID, sol.AT, sol.MP, AT, MP)
		return &pb.Respuesta{Rpta: 0}, nil
	}
	if (sol.MP > MP) {
		fmt.Printf("Recepcion de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA -- AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", sol.ID, sol.AT, sol.MP, AT, MP)
		return &pb.Respuesta{Rpta: 0}, nil
	}
	AT -= sol.AT
	MP -= sol.MP
	fmt.Printf("Recepcion de solicitud desde equipo %d, %d AT y %d MP -- APROBADA -- AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", sol.ID, sol.AT, sol.MP, AT, MP)
	return &pb.Respuesta{Rpta: 1}, nil
}

// Funcion que produce misiles cada 1 segundo
func producirMisiles(wg *sync.WaitGroup) {
    defer wg.Done()
	
    for 
	{
		// Produce AT, si se pasa queda en su maximo valor
		if (AT < maxAT) {
			AT += 10
		} 
		if (AT >= maxAT) {
			AT = maxAT
		}
		
		// Produce MP, si se pasa queda en su maximo valor
		if (MP < maxMP) {
			MP += 5
		} 
		if (MP >= maxMP) {
			MP = maxMP
		}

		// Pausa de 1 segundo
        time.Sleep(1 * time.Second)
    }
}

// Esucha las llamadas de gRPC en un proceso separado
func escucharLlamadas(wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Abre y escucha al puerto
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// Se crea un servidor gRPC
	serverRegister := grpc.NewServer()
	service := &myGemaEstrategicaServer{}
	pb.RegisterGemaEstrategicaServer(serverRegister, service)
	
	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}

func main() {
	var wg sync.WaitGroup

    wg.Add(2)
	
	// Escucha al servidor y produce misiles simultaneamente
	go escucharLlamadas(&wg)
	go producirMisiles(&wg)

	// Esperamos a que ambas GoRoutines finalicen
	wg.Wait()
}
