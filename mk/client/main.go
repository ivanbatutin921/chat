package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "root/mk/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure()) // Подключение к серверу
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
			in, err := stream.Recv() // Чтение ответа от сервера
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
			}
			log.Printf("Received echo: %v", in.GetMessage())
		}
	}()

	for {
		if err := stream.Send(&pb.LiveChatData{Message: "Hello, server!"}); err != nil { // Отправка сообщения серверу
			log.Fatalf("Failed to send a message: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}