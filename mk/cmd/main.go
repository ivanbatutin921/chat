package main

import (
	"fmt"
	"log"
	"net"

	db "root/mk/internal/database"
	"root/mk/internal/server"
	pb "root/mk/proto"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()
	//db.Migration()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLiveChatServer(s, &server.Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("server stop")

}
