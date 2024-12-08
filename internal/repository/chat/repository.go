package chat

import (
	"context"
	def "github.com/sadstill/chat-server/internal/repository"
	repoModel "github.com/sadstill/chat-server/internal/repository/chat/model"
	"sync"
)

var _ def.ChatRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.Message
	m    sync.RWMutex
}

func NewRepository() def.ChatRepository {
	return &repository{
		data: make(map[string]*repoModel.Message),
	}
}

func (r *repository) Create(ctx context.Context, usernames []string) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, chatId int64) {
	//TODO implement me
	panic("implement me")
}
