package main

import (
	"context"
	"fmt"
	desc "github.com/sadstill/chat-server/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) SendMessage(ctx context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received request: %+v", in)
	_ = ctx
	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %d", grpcPort)

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
