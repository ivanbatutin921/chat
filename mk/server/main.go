package main

import (
	"log"
	"net"
	"sync"
	"sync/atomic"

	pb "root/mk/proto"

	"google.golang.org/grpc"
)

var clientIdCounter int32

type Server struct {
	pb.UnimplementedLiveChatServer
	clientStreams sync.Map // Использование sync.Map для хранения потоков
}

func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) error {
	clientID := atomic.AddInt32(&clientIdCounter, 1)

	s.clientStreams.Store(clientID, stream)

	defer s.clientStreams.Delete(clientID)

	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("Client %v disconnected: %v", clientID, err)
			return err
		}

		s.clientStreams.Range(func(key, value any) bool {
			if key != clientID { // Отправка сообщения всем, кроме источника
				otherStream, ok := value.(pb.LiveChat_ChatStreamServer)
				if ok {
					if sendErr := otherStream.Send(in); sendErr != nil {
						log.Printf("Failed to send message to client %v: %v", key, sendErr)
					}
				}
			}
			return true
		})
	}
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLiveChatServer(s, &Server{})

	log.Println("Server listening on port :9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
