package chat

import desc "github.com/sadstill/chat-server/pkg/chat/v1"

type Implementation struct {
	desc.UnimplementedChatV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
