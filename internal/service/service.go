package service

import (
	"context"
	"github.com/sadstill/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, id int64)
	Send(ctx context.Context, message *model.Message)
}
