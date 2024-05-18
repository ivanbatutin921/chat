package main

import (
	"bufio"
	"context"
	"log"
	"os"
	pb "root/mk/proto"

	"google.golang.org/grpc"
)

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

	// Обработка входящих сообщений
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive a message : %v", err)
			}
			log.Printf("Received message: %s", in.GetMessage())
		}
	}()

	// Чтение сообщений с консоли и отправка их
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if err := stream.Send(&pb.LiveChatData{Message: message}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}
