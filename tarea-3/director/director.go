package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"math/rand"
	"net"
	"strconv"
)

var playerCounter int
var x int
var y int
var h int
var p int
var hp int

type server struct {
	pb.UnimplementedDirectorCommunicationServer
}

func (s *server) ConfirmReady(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	playerCounter += 1
	if playerCounter > 8 {
		fmt.Println("All players are ready")
	}
	return &pb.Response{Res: "ok"}, nil
}

func (s *server) PickGun(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	SendDecisionToNameNode(in.ClientIp, in.Msg, "1")
	if in.Msg == "1" {
		if rand.Intn(x) > 100 {
			killPlayer(in.ClientIp)
			return &pb.Response{Res: "dead"}, nil
		}
	} else if in.Msg == "2" {
		if rand.Intn(y-x) > 100 {
			killPlayer(in.ClientIp)
			return &pb.Response{Res: "dead"}, nil
		}
	} else {
		if rand.Intn(100-y) > 100 {
			killPlayer(in.ClientIp)
			return &pb.Response{Res: "dead"}, nil
		}
	}
	return &pb.Response{Res: "next floor"}, nil
}

func (s *server) PickHallway(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	SendDecisionToNameNode(in.ClientIp, in.Msg, "2")
	if in.Msg == "A" {
		if h == 1 {
			killPlayer(in.ClientIp)
			return &pb.Response{Res: "dead"}, nil
		}
	} else if h == 0 {
		killPlayer(in.ClientIp)
		return &pb.Response{Res: "dead"}, nil
	}
	return &pb.Response{Res: "next floor"}, nil
}

func (s *server) BossBattle(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	SendDecisionToNameNode(in.ClientIp, in.Msg, "3")
	number, _ := strconv.Atoi(in.Msg)
	if number == p {
		hp = -1
	}
	if hp == 0 {
		return &pb.Response{Res: "boss dead"}, nil
	}
	return &pb.Response{Res: "boss not dead"}, nil
}

func (s *server) AskForAccountBalance(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	// TODO:
	return &pb.Response{Res: ""}, nil
}

func main() {
	playerCounter = 0
	fmt.Print("Starting server...\n")
	PlayGame()
	StartServer()
}

func StartServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al crear servidor GRPC en la dirección %s: %s", ":50051", err)
	}
	grpcServer := grpc.NewServer()
	service := &server{}
	pb.RegisterDirectorCommunicationServer(grpcServer, service)
	log.Println("Server listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func PlayGame() {
	piso1()
	piso2()
	piso3()
}

func piso1() {
	x = rand.Intn(100)
	y = rand.Intn(100-x) + x
}

func piso2() {
	h = rand.Intn(2)
}

func piso3() {
	p = rand.Intn(15) + 1
	hp = 2
}

func SendDecisionToNameNode(ip string, decision string, floor string) {

	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewNodesCommunicationClient(conn)

	client.NotifyDecision(context.Background(), &pb.Decision{Option: decision, ClientIp: ip, Floor: floor})

	conn.Close()
}

func killPlayer(ip string) {
	// TODO:
	fmt.Println("Player " + ip + " has died.")
}
