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

type NotificationDialogResponse struct {
	Notifications    []NotificationDialogResponse `json:"notifications"`
	UnreadNotiAmount int                          `json:"unread_noti_amount"`
}

type NotificationResponse struct {
	NotificationId string    `json:"notification_id"`
	ActorUsername  string    `json:"actor_username"`
	ActorAvatar    string    `json:"actor_avatar"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	IsRead         bool      `json:"is_read"`
}
