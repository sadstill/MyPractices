package chat

import (
	"context"
	desc "github.com/sadstill/chat-server/pkg/chat/v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, req.Usernames)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		ChatId: id,
	}, nil
}
