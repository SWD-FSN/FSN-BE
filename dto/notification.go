package dto

import "time"

type CreateNotiRequest struct {
	ActorId    string `json:"actor_id"`
	ObjectId   string `json:"object_id"` // Post, cmt, follow của tài khoản được tác động
	ObjectType string `json:"object_type"`
	Action     string `json:"action"` // Like, follow
}

type GetNotiOnActionRequest struct {
	ActorId    string    `json:"actor_id"`
	ObjectId   string    `json:"object_id"` // Post, cmt, follow của tài khoản được tác động
	ObjectType string    `json:"object_type"`
	Action     string    `json:"action"` // Like, follow
	CreatedAt  time.Time `json:"created_at"`
}

type NotificationResponse struct {
	NotificationId string `json:"notification_id"`
}
