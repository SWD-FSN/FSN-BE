package service

import (
	"context"
	"social_network/dto"
)

type IConversationService interface {
	GetAllConversations(ctx context.Context) (*[]dto.ConversationResponse, error)
	GetConversationsFromUser(id string, ctx context.Context) *[]dto.ConversationUIResponse
	GetConversationFromUser(actorId, conversationId string, ctx context.Context) (*dto.InternalConversationUIResponse, error)
	GetConversationsByKeywordFromUser(id, keyword string, ctx context.Context) *[]dto.ConversationSearchBarResponse
	GetMessagesInChatByKeyword(req dto.SearchMessagesInChatRequest, ctx context.Context) (*dto.SearchMessagesInChatResponse, error)
	CreateConversation(req dto.CreateConversationRequest, ctx context.Context) (*dto.ConversationUIResponse, error)
	EditGroupChatProperty(req dto.EditGroupChatPropRequest, ctx context.Context) error
	DissovelGroupConversation(actorId, conversationId string, ctx context.Context) error
	LeaveGroupConversation(memberId, conversationId string, ctx context.Context) error
}
