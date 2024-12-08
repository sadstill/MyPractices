package main

import (
	"fmt"
	"github.com/sadstill/chat-server/internal/delivery/chat"
	chatRepo "github.com/sadstill/chat-server/internal/repository/chat"
	chatService "github.com/sadstill/chat-server/internal/service/chat"
	"log"
	"net"

	desc "github.com/sadstill/chat-server/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatRepository := chatRepo.NewRepository()
	chatSrv := chatService.NewService(chatRepository)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, chat.NewImplementation(chatSrv))

	log.Printf("server listening at %d", grpcPort)

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
