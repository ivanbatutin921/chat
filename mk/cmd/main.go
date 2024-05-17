package main

import (
	"log"
	"net"

	pb "root/mk/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
	pb.UnimplementedLiveChatServer
}

func (s *server) SendMessage(stream pb.LiveChat_ChatStreamClient) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", msg)
		// Дополнительная обработка полученного сообщения
	}
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLiveChatServer(s, &server{})
	reflection.Register(s)

	log.Println("Server listening on port :9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
