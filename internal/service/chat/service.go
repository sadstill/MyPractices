package chat

import (
	"github.com/sadstill/chat-server/internal/repository"
	def "github.com/sadstill/chat-server/internal/service"
)

var _ def.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
}

func NewService(
	chatRepository repository.ChatRepository,
) def.ChatService {
	return &service{
		chatRepository: chatRepository,
	}
}
