package businessobject

import "time"

type Notification struct {
	NotificationId string    `json:"notification_id"`
	ActorId        string    `json:"actor_id"` // Người thực hiện hành động like, follow, ...
	TargetUserId   string    `json:"target_user_id"`
	PostId         string    `json:"post_id"`
	CommentId      string    `json:"comment_id"`
	Action         string    `json:"action"` // Like, follow, comment, rep comment
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}

func GetNotificationTable() string {
	return "notifications"
}
