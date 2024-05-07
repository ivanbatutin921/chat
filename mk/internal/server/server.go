package server

import (
	pb "root/mk/proto"
)

type Server struct {
	pb.UnimplementedLiveChatServer
}

func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) (*pb.LiveChatData, error) {
	for {
		var data pb.LiveChatData
		err := stream.RecvMsg(&data)
		if err != nil {
			return nil, err
		}

		if data.Message == "" {
			return &pb.LiveChatData{Message: "{\"ошибка\":\"сообщение не может быть пустым\"}"}, nil

		}

		err = stream.Send(&pb.LiveChatData{Message: data.Message})
		if err != nil {
			return nil, err
		}
		///return &pb.LiveChatData{Message: data.Message},nil
	}
}

// func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
// 	err := db.DB.DB.Create(&model.User{Name: req.Name, Email: req.Email})
// 	if err.Error != nil {
// 		log.Fatal(err.Error)
// 	}
// 	return &pb.User{
// 			Name:  req.Name,
// 			Email: req.Email},
// 		nil
// }
