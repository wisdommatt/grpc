package main

import (
	"log"
	"net"

	"github.com/wisdommatt/grpc/chat"
	"github.com/wisdommatt/grpc/chat/chat"

	"google.golang.org/grpc"
)

func main() {
	// personGrpc()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	chatServer := &chat.Server{}
	chat.RegisterChatServiceService(grpcServer, chatServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
