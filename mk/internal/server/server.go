package server

import (
<<<<<<< HEAD
	"fmt"
	"log"
=======
	"errors"

	bd "root/mk/internal/database"
	model "root/mk/internal/model"
>>>>>>> bae133bcdf1c7798d60b3a10fc7b1feceee3d39d
	pb "root/mk/proto"
	"time"
)

type Server struct {
	pb.UnimplementedLiveChatServer
}

<<<<<<< HEAD
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
=======
func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) error {
>>>>>>> bae133bcdf1c7798d60b3a10fc7b1feceee3d39d
	for {
		msg, err := stream.Recv()
		if err != nil {
<<<<<<< HEAD
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Received message from server: %s", msg.GetMessage())
	}

	<-done
	return nil
=======
			return err
		}

		if data.Message == "" {
			return errors.New("empty message")
		}

		err = stream.Send(&pb.LiveChatData{Message: data.Message})
		if err != nil {
			return err
		}
		//260~270 ms это ммм(((
		bd.DB.DB.Create(&model.Message{Message: data.Message})
	}
>>>>>>> bae133bcdf1c7798d60b3a10fc7b1feceee3d39d
}
