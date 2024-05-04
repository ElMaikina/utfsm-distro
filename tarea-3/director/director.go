package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/proto"
	"context"
"net"
)

type myDirectorCom
	mustEmbedUnimplementedDirectorComServer
}

func (s myDirectorComServer ) SendMessage(context.Context, *Message) (*Response, error) {
	return Response{rpta: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myDirectorComServer{}
	RegisterDirectorComServer(serverRegister, service)

	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
