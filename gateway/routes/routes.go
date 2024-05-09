package routes

import (
	pb "root/mk/proto"

	"github.com/gofiber/fiber/v2"
)

type ChatServiceHandler struct {
	pb.UnimplementedLiveChatServer
	mk pb.LiveChatClient
}

func ServiceHandler(mk pb.LiveChatClient) *ChatServiceHandler {
	return &ChatServiceHandler{mk: mk}
}


func (c *ChatServiceHandler) HandleChatStream(ctx *fiber.Ctx) error {
	var data pb.LiveChatData
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	
	
	
	stream, err := c.mk.ChatStream()
	if err != nil {
		return err
	}
	if err := stream.Send(&data); err != nil {
		return err
	}
	receivedData, err := stream.Recv()
	if err != nil {
		return err
	}

	return ctx.JSON(receivedData)
}
