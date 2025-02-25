package businessobject

import "time"

type Notification struct {
	NotificationId string    `json:"notification_id"`
	ActorId        string    `json:"actor_id"`  // Người thực hiện hành động like, follow, ...
	ObjectId       string    `json:"object_id"` // Post, cmt, follow của tài khoản được tác động
	ObjectType     string    `json:"object_type"`
	Action         string    `json:"action"` // Like, follow
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}

func GetNotificationTable() string {
	return "notification"
}
