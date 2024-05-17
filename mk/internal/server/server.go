package server

import (
	"fmt"
	"log"
	pb "root/mk/proto"
	"time"
)

type Server struct {
	pb.UnimplementedLiveChatServer
}

func (s *Server) ClientStream(stream pb.LiveChat_ChatStreamClient) error {
	done := make(chan bool)

	// Send messages to the server
	go func() {
		for i := 0; i < 5; i++ {
			message := fmt.Sprintf("Message %d from client", i+1)
			err := stream.Send(&pb.LiveChatData{Message: message})
			if err != nil {
				log.Fatalf("error sending message: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		done <- true
	}()

	// Receive messages from the server
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Received message from server: %s", msg.GetMessage())
	}

	<-done
	return nil
}
