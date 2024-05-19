package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
)

func main() {
	ip := "1"

	conn, err := grpc.Dial("200.1.22.237:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewDirectorCommunicationClient(conn)

	fmt.Println("Presiona enter para comenzar")
	fmt.Scanln()
	response, _ := client.ConfirmReady(context.Background(), &pb.Message{Msg: "1", ClientIp: ip})

	// Da la opcion al jugador de elegir una de tres armas
	fmt.Println("Selecciona tu arma (elige por numero)")
	fmt.Println(" 1) Escopeta")
	fmt.Println(" 2) Rifle automatico")
	fmt.Println(" 3) Punos electricos")

	var gun string
	fmt.Scanln(&gun)

	response, _ = client.PickGun(context.Background(), &pb.Message{Msg: gun, ClientIp: ip})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}

	// Elige uno de dos pasillos
	fmt.Println("Selecciona un pasillo (A o B)")
	var hallway_str string
	fmt.Scanln(&hallway_str)
	response, _ = client.PickHallway(context.Background(), &pb.Message{Msg: hallway_str, ClientIp: ip})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}

	// Se da la opcion de elegir uno de quince pisos
	fmt.Println("Selecciona un numero (del 1 al 15)")
	var number string
	fmt.Scanln(&number)
	response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: ip})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}

	// Se da la opcion de elegir uno de quince pisos
	fmt.Println("Selecciona un numero (del 1 al 15)")
	fmt.Scanln(&number)
	response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: ip})
	fmt.Println(response.Res)

	conn.Close()
}
