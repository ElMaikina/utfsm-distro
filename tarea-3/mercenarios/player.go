package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "main/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer conn.Close()

	// Crea un nuevo cliente gRPC
	client := pb.NewCheckStatusClient(conn)

	for {
		response, err := client.SolicitarM(context.Background(), request)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		// Imprime la respuesta del servidor
		if response.Rpta == 1 {
			fmt.Printf("Positivo: Solicitando %d AT y %d MP ; Resolucion: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicacion\n", AT, MP)
		}
		if response.Rpta == 0 {
			fmt.Printf("Negativo: Solicitando %d AT y %d MP ; Resolucion: -- DENEGADA -- ; Reintentando en 3 segs...\n", AT, MP)
		}
		time.Sleep(3 * time.Second)
	}
}
