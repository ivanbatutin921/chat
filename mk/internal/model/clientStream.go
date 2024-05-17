package model

import (
	pb "root/mk/proto"
)

type ClientStream struct {
	Id int32
	Stream pb.LiveChat_ChatStreamServer
}