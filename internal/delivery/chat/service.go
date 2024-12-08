package chat

import (
	"github.com/sadstill/chat-server/internal/service"
	desc "github.com/sadstill/chat-server/pkg/chat/v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
