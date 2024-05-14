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

type server struct{}

var playerCounter int
var x int
var y int
var h int
var p int
var hp int

func (s *server) ConfirmReady(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	playerCounter += 1
	if playerCounter > 8 {
		playGame()
	}
	return &pb.Response{res: "ok"}, nil
}

func (s *server) PickGun(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	if in.msg == '1' {
		if rand.Intn(x) > 100 {
			killPlayer(in.client_ip)
			return &pb.Response{res: "dead"}, nil
		}
	} else if in.msg == '2' {
		if rand.Intn(y-x) > 100 {
			killPlayer(in.client_ip)
			return &pb.Response{res: "dead"}, nil
		}
	} else {
		if rand.Intn(100-y) > 100 {
			killPlayer(in.client_ip)
			return &pb.Response{res: "dead"}, nil
		}
	}
	return &pb.Response{res: "next floor"}, nil
}

func (s *server) PickHallway(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	if in.msg == 'A' {
		if h == 1 {
			killPlayer(in.client_ip)
			return &pb.Response{res: "dead"}, nil
		}
	} else if h == 0 {
		killPlayer(in.client_ip)
		return &pb.Response{res: "dead"}, nil
	}
	return &pb.Response{res: "next floor"}, nil
}

func (s *server) BossBattle(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	number, _ := strconv.Atoi(in.msg)
	if number == p {
		hp = -1
	}
	if hp == 0 {
		return &pb.Response{res: "boss dead"}, nil
	}
	return &pb.Response{res: "boss not dead"}, nil
}

func (s *server) AskForAccountBalance(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	// TODO:
	return &pb.Response{res: ""}, nil
}

func main() {
	playerCounter = 0
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al crear servidor GRPC en la dirección %s: %s", ":50051", err)
	}
	grpcServer := grpc.newServer()
	pb.RegisterDirectorCommunicationService(grpcServer, &server{})
	log.Println("Server listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func playGame() {
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
	p = rand.Intn(15)
	hp = 2
}

func killPlayer(ip int) {
	// TODO:
}
