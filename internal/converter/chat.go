package converter

import (
	"github.com/sadstill/chat-server/internal/model"
	desc "github.com/sadstill/chat-server/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToMessageFromService(message *model.Message) *desc.Message {
	var timestamp *timestamppb.Timestamp
	if message.Timestamp != nil {
		timestamp = timestamppb.New(*message.Timestamp)
	}

	return &desc.Message{
		ChatId:    message.ChatId,
		From:      message.From,
		Text:      message.Text,
		Timestamp: timestamp,
	}
}
