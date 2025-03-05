package repo

import (
	"context"
	business_object "social_network/business_object"
)

type IConversationRepo interface {
	GetAllConversations(ctx context.Context) (*[]business_object.Conversation, error)
	GetConversationsFromUser(id string, ctx context.Context) (*[]business_object.Conversation, error)
	GetConversationOfTwoUsers(userId1, userId2 string, ctx context.Context) (*business_object.Conversation, error)
	GetConversationsByKeyword(id, keyword string, ctx context.Context) (*[]business_object.Conversation, error)
	GetConversation(id string, ctx context.Context) (*business_object.Conversation, error)
	CreateConversation(conversation business_object.Conversation, ctx context.Context) error
	DissovelGroupConversation(id string, ctx context.Context) error
	UpdateConversation(conversation business_object.Conversation, ctx context.Context) error
	//RemoveConversation(id string, ctx context.Context) error
}
