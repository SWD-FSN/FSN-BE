package repo

import (
	"context"
	business_object "social_network/business_object"
)

type IMessageRepo interface {
	GetAllMessages(ctx context.Context) (*[]business_object.Message, error)
	GetMessagesFromConversation(id string, ctx context.Context) (*[]business_object.Message, error)
	GetMessagesFromConversationByKeyword(id, keyword string, ctx context.Context) (*[]business_object.Message, error)
	GetMessagesFromUser(id string, ctx context.Context) (*[]business_object.Message, error)
	GetMessage(id string, ctx context.Context) (*business_object.Message, error)
	CreateMessage(message business_object.Message, ctx context.Context) error
}
