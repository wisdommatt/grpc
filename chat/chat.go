package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server grpc chat server
type Server struct {
}

// SayHello says hello to grpc client
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from the server !"}, nil
}
