package repository

import (
	"context"
)

type ChatRepository interface {
	Create(ctx context.Context, usernames []string)
	Delete(ctx context.Context, chatId int64)
}
