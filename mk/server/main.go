package main

import (
	"log"
	"net"
	"sync"
	"sync/atomic"

	"root/mk/internal/model"
	pb "root/mk/proto"

	"google.golang.org/grpc"
)

var clientIdCounter int32

type Server struct {
	pb.UnimplementedLiveChatServer
	clientStreams sync.Map // Использование sync.Map для хранения потоков
}

func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) error {
    // Генерируем уникальный ID для нового клиента
    NewClientId := atomic.AddInt32(&clientIdCounter,1) // Функция для генерации ID не показана

    // Регистрируем поток нового клиента
    s.clientStreams.Store(NewClientId, model.ClientStream{Id: int32(NewClientId), Stream: stream})

    defer s.clientStreams.Delete(NewClientId) // Удаляем поток из хранилища при завершении

    for {
        in, err := stream.Recv()
        if err != nil {
            log.Printf("client %d disconnected: %v", NewClientId, err)
            return err
        }
        // Перебираем все потоки и отправляем им полученное сообщение
        s.clientStreams.Range(func(key, value any) bool {
            clientStream, ok := value.(model.ClientStream)
            if ok && clientStream.Id != int32(NewClientId) { // Не отправляем сообщение обратно отправителю
                if err := clientStream.Stream.Send(in); err != nil {
                    log.Printf("Failed to send message to client %d: %v", clientStream.Id, err)
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
