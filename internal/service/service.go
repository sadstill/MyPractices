package service

import "context"

type ChatService interface {
	Create(ctx context.Context)
	Get(ctx context.Context)
	Delete(ctx context.Context)
}
