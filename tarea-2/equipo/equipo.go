package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"math/rand"
	"time"
)

func main() {
	// Establece la conexión al servidor gRPC
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Crea un nuevo cliente gRPC
	client := pb.NewGreeterClient(conn)

	time.Sleep(10 * time.Second)

	for {
		// Define el mensaje a enviar al servidor
		AT := int32(20 + rand.Intn(10))
		MP := int32(10 + rand.Intn(5))
		request := &pb.Solicitud{AT: AT, MP: MP}

		// Llama al método remoto del servidor gRPC
		response, err := client.solicitarM(context.Background(), request)
		if err != nil {
			log.Fatalf("Error calling SolicitarM: %v", err)
		}

		// Imprime la respuesta del servidor
		log.Printf("Response from server: %s", response.Message)

		if response.Message == 1 {
			break
		}

	}
	time.Sleep(3 * time.Second)
}
