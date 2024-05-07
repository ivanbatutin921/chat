package main

import (
	"context"
	"log"

	pb "root/mk/proto"

	"google.golang.org/grpc"
)

type client struct {
	pb.UnimplementedLiveChatServer
}

func (s *client) ClientStream(stream pb.LiveChat_ChatStreamClient) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		err = stream.Send(&pb.LiveChatData{Message: req.Message})
		if err != nil {
			return err
		}
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewLiveChatClient(conn)

	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("error receiving message: %v", err)
	}

	go func() {

		err := stream.Send(&pb.LiveChatData{Message: "Hello, server!"})
		if err != nil {
			log.Fatalf("error sending message: %v", err)
		}

	}()

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Received message: %s", msg.GetMessage())
	}
}
