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

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewDirectorCommunicationClient(conn)

	fmt.Println("Presiona enter para comenzar")
	fmt.Scanln()
	response, _ := client.ConfirmReady(context.Background(), &pb.Message{Msg: "1", ClientIp: ip})

	fmt.Println("Selecciona tu arma")
	var gun string
	fmt.Scanln(&gun)

	response, _ = client.PickGun(context.Background(), &pb.Message{Msg: gun, ClientIp: ip})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}

	fmt.Println("Selecciona un pasillo")
	var hallway_str string
	fmt.Scanln(&hallway_str)
	response, _ = client.PickHallway(context.Background(), &pb.Message{Msg: hallway_str, ClientIp: ip})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}

	fmt.Println("Selecciona un numero del 1 al 15")
	var number string
	fmt.Scanln(&number)
	response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: ip})
	fmt.Println(response.Res)

	if response.Res == "boss not dead" {
		fmt.Println("Selecciona un numero del 1 al 15")
		var number string
		fmt.Scanln(&number)
		response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: ip})
	}

	fmt.Println(response.Res)

	conn.Close()
}
