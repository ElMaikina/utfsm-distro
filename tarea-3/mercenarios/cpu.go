package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"math/rand"
	"strconv"
)

func main() {
	// Direccion IP de la maquina
	ipAddress := "200.1.22.237"
	port := "50051"
	ip := ipAddress+":"+port

	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewDirectorCommunicationClient(conn)

	response, _ := client.ConfirmReady(context.Background(), &pb.Message{Msg: "1", ClientIp: ip})
	if response.Res == "dead" {
		return
	}
	gun_number := rand.Intn(3) + 1
	gun := strconv.Itoa(gun_number)
	response, _ = client.PickGun(context.Background(), &pb.Message{Msg: gun, ClientIp: ip})
	if response.Res == "dead" {
		return
	}
	
	// Elige uno de dos pasillos
	hallway := rand.Intn(2)
	hallway_str := ""
	if hallway == 0 {
		hallway_str = "A"
	} else {
		hallway_str = "B"
	}
	response, _ = client.PickHallway(context.Background(), &pb.Message{Msg: hallway_str, ClientIp: gun})
	if response.Res == "dead" {
		return
	}
	
	// Se da la opcion de elegir uno de quince pisos
	number_n := rand.Intn(15) + 1
	number := strconv.Itoa(number_n)
	response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: gun})
	fmt.Println(response.Res)
	if response.Res == "dead" {
		return
	}
	
	// Se da la opcion de elegir uno de quince pisos
	number_n = rand.Intn(15) + 1
	number = strconv.Itoa(number_n)
	response, _ = client.BossBattle(context.Background(), &pb.Message{Msg: number, ClientIp: gun})
	fmt.Println(response.Res)

	conn.Close()
}
