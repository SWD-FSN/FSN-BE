package dto

import "time"

// Response
type ConversationResponse struct {
	ConversationId     string    `json:"conversation_id"`
	ConversationAvatar string    `json:"conversation_avatar"`
	ConversationName   []string  `json:"conversation_name"` // Tên đoạn chat. Nếu là đoạn chat giữa 2 user thì tên sẽ có 2 hiện tùy bên user dựa trên username. Nếu là group chat sẽ có 1 tên chung
	HostId             *string   `json:"host_id"`           // Trưởng nhóm nếu đây là group chat
	Members            []string  `json:"members"`           // Nếu đoạn chat giữa 2 người sẽ mặc định 2
	IsGroup            bool      `json:"is_group"`          // Có phải là group chat hay ko
	IsDelete           *bool     `json:"is_delete"`         // Nếu là group chat thì trưởng nhóm có quyền giải tán nhóm
	CreatedAt          time.Time `json:"created_at"`
}

type ConversationUIResponse struct {
	ConversationId     string    `json:"conversation_id"`
	ConversationAvatar string    `json:"conversation_avatar"`
	ConversationName   string    `json:"conversation_name"`
	LatestMessageId    string    `json:"latest_message_id"`
	MessageContent     string    `json:"message_contet"`
	SentAt             time.Time `json:"sent_at"`
	SenderId           string    `json:"sender_id"`
	SenderUsername     string    `json:"sender_username"`
}

type InternalConversationUIResponse struct {
	ConversationId     string               `json:"conversation_id"`
	ConversationAvatar string               `json:"conversation_avatar"`
	ConversationName   string               `json:"conversation_name"`
	ActorMessages      *[]MessageUIResponse `json:"actor_messages"`
	MemberMessages     *[]MessageUIResponse `json:"member_messages"`
}

type InternalConversationUIResponseV2 struct {
	ConversationId     string               `json:"conversation_id"`
	ConversationAvatar string               `json:"conversation_avatar"`
	ConversationName   string               `json:"conversation_name"`
	RequestUserid      string               `json:"request_user_id"`
	Messages           *[]MessageUIResponse `json:"messages"`
}

type ConversationSearchBarResponse struct {
	ConversationId     string `json:"conversation_id"`
	ConversationAvatar string `json:"conversation_avatar"`
	ConversationName   string `json:"conversation_name"`
}

type SearchMessagesInChatResponse struct {
	ConversationId     string               `json:"conversation_id"`
	ConversationName   string               `json:"conversation_name"`
	ConversationAvatar string               `json:"conversation_avatar"`
	Messages           *[]MessageUIResponse `json:"messages"`
}

/* -------------------------------------------------- */

// Request
type SearchMessagesInChatRequest struct {
	ConversationId string `json:"conversation_id"`
	ActorId        string `json:"actor_id"`
	Keyword        string `json:"keyword"`
}

type CreateConversationRequest struct {
	ActorId string   `json:"actor_id"`
	Members []string `json:"members"`
	IsGroup bool     `json:"is_group"`
}

type EditGroupChatPropRequest struct {
	ConversationId string `json:"conversation_id"`
	ActorId        string `json:"actor_id"`
	Property       string `json:"property"`
	Value          string `json:"value"`
}
