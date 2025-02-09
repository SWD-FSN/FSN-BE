package businessobject

import "time"

type Notification struct {
	NotificationId string    `json:"notification_id"`
	ActionId       string    `json:"action_id"`
	AccountId      string    `json:"account_id"` // Tài khoản được tác động đến
	ObjectId       string    `json:"object_id"`  // Post, cmt, follow của tài khoản được tác động
	CreatedAt      time.Time `json:"created_at"`
	IsRead         bool      `json:"is_read"`
	Status         bool      `json:"status"`
}

func GetNotificationTable() string {
	return "notification"
}
