package dto

import "time"

type WSSendMessageRequest struct {
	UserId             string    `json:"user_id"`
	UserAvatar         string    `json:"user_avatar"`
	Content            any       `json:"content"`
	ContentType        string    `json:"content_type"`
	CreatedAt          time.Time `json:"created_at"`
	ContentContainerId *string   `json:"content_container_id"`
}
