package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"math/rand"
	"net"
	"strconv"
)

func main() {
	ip := 1

	conn, err := grpc.Dial(":50051")
	if err != nil {
		log.Fatal("Error:" + err)
	}
	defer conn.Close()

	client := pb.NewDirectorCommunicationClient(conn)

	// for i := 0; i < len(array); i++ {
	// 	fmt.Println("Element at index", i, ":", array[i])
	// }

	response, _ := client.ConfirmReady(context.Background(), &pb.Message{msg: '1', client_ip: ip})
	if response == "dead" {
		return
	}
	gun_number := rand.Intn(3) + 1
	gun := strconv.Itoa(gun_number)
	response, _ = client.PickGun(context.Background(), &pb.Message{msg: gun, client_ip: ip})
	if response == "dead" {
		return
	}
	hallway := rand.Intn(2)
	if hallway == 0 {
		hallway = 'A'
	} else {
		hallway = 'B'
	}
	response, _ = client.PickHallway(context.Background(), &pb.Message{msg: hallway, client_ip: gun})
	if response == "dead" {
		return
	}
	number_n := rand.Intn(15) + 1
	number := strconv.Itoa(number_n)
	response, _ = client.BossBattle(context.Background(), &pb.Message{msg: number, client_ip: gun})

	conn.Close()
}
