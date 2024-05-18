package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	pb "root/mk/proto"

	"google.golang.org/grpc"
)

var clientCount int32

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLiveChatClient(conn)
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv() // Получение сообщений от сервера
			if err != nil {
				log.Printf("Failed to receive: %v", err)
				return
			}
			log.Printf("Received message: %v", in.GetMessage())
		}
	}()

	for {
		// Отправка пользовательского сообщения
		clientCount := atomic.AddInt32(&clientCount, 1)
		if err := stream.Send(&pb.LiveChatData{Message: fmt.Sprintf("Client %d: Hello from client!", clientCount)}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}
