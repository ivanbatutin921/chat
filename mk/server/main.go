package main

import (
	"log"
	"net"
	"sync"

	model "root/mk/internal/model"
	pb "root/mk/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedLiveChatServer
	clientStreams sync.Map
}

func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) error {
	//получаем id откудато
	NewClientId := 1

	//регестрируем поток клиента
	s.clientStreams.Store(NewClientId, model.ClientStream{Id: int32(NewClientId), Stream: stream})

	for {
		in, err := stream.Recv()
		if err != nil {
			s.clientStreams.Delete(NewClientId)
			return err
		}
		s.clientStreams.Range(func(key, value any) bool {
			clientStreams, ok := value.(model.ClientStream)
			if ok && clientStreams.Id != int32(NewClientId) {
				if err := clientStreams.Stream.Send(in); err != nil {
					log.Println("Failed to send message:", err)
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
	reflection.Register(s)

	log.Println("Server listening on port :9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
