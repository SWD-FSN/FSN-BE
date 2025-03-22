package dto

import "time"

type MessageUIResponse struct {
	MessageId     string    `json:"message_id"`
	ConvesationId string    `json:"conversation_id"`
	AuthorId      string    `json:"author_id"`
	AuthorAvatar  string    `json:"author_avatar"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateMessageRequest struct {
	AuthorId       string `json:"author_id" validate:"required"`
	ConversationId string `json:"conversation_id" validate:"required"`
	Content        string `json:"content" validate:"required"`
}
