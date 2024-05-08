package server

import (
	"errors"

	pb "root/mk/proto"
	bd "root/mk/internal/database"
	model "root/mk/internal/model"
)

type Server struct {
	pb.UnimplementedLiveChatServer
}

func (s *Server) ChatStream(stream pb.LiveChat_ChatStreamServer) error {
	for {
		var data pb.LiveChatData
		err := stream.RecvMsg(&data)
		if err != nil {
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
}

