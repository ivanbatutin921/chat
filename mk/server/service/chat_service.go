package server

import (
	pb "root/mk/proto"
)

type Server struct {
	pb.UnimplementedLiveChatServer
}
