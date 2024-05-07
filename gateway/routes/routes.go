package routes

import (
	"context"
	"log"
	"net/http"
	pb "root/mk/proto"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ChatServiceHandler struct {
	pb.UnimplementedLiveChatServer
	mk pb.LiveChatServer
}

func ServiceHandler(mk pb.LiveChatServer) *ChatServiceHandler {
	return &ChatServiceHandler{mk: mk}
}

func (c *ChatServiceHandler) GetMessage(f *fiber.Ctx) error {
	body := pb.LiveChatData{}
	if body.Message == "" {
		return f.JSON(fiber.Map{"status": "error", "message": "тело запроса путое"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

}

func(c *ChatServiceHandler) handleChatStream(w http.ResponseWriter, r *http.Request) {
	// Создание потока для отправки данных обратно клиенту
	stream := pb.NewLiveChatClient(w, r)

	// Вызов метода ChatStream
	err := c.mk.ChatStream(stream)
	if err != nil {
		// Обработка ошибки
		log.Printf("Error in ChatStream: %v", err)
		return
	}
}
