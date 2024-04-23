package main

import (
	"context"
	"google.golang.org/grpc"
	"math/rand"
	"log"
	"fmt"
	"time"
	pb "main/proto"
)

func main() {
	// Establece la conexión al servidor gRPC
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer conn.Close()

	// Crea un nuevo cliente gRPC
	client := pb.NewGemaEstrategicaClient(conn)

	time.Sleep(10 * time.Second)

	for {
		// Define el mensaje a enviar al servidor
		AT := int32(20 + rand.Intn(10))
		MP := int32(10 + rand.Intn(5))
		request := &pb.Solicitud{AT: AT, MP: MP}

		// Llama al método remoto del servidor gRPC
		response, err := client.SolicitarM(context.Background(), request)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		// Imprime la respuesta del servidor
		if (response.Rpta == 1) {
			fmt.Printf("Positivo: Solicitando %d AT y %d MP ; Resolucion: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicacion\n", AT, MP)
		}
		if (response.Rpta == 0) {
			fmt.Printf("Negativo: Solicitando %d AT y %d MP ; Resolucion: -- DENEGADA -- ; Reintentando en 3 segs...\n", AT, MP)
		}
		time.Sleep(3 * time.Second)
	}
}
