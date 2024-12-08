package repository

import (
	"context"
	"github.com/sadstill/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, message *model.Message)
	Delete(ctx context.Context)
	Send(ctx context.Context)
}
