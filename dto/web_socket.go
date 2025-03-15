package dto

type WSSendMessageRequest struct {
	Id          string `json:"id"`
	Content     any    `json:"content"`
	ContentType string `json:"content_type"`
}
