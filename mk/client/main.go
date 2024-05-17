package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "root/mk/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLiveChatClient(conn)
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	for {
		msg := &pb.Message{Text: "Hello, server!"}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		time.Sleep(time.Second) // Отправляем сообщение каждую секунду
	}
}