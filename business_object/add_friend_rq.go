package businessobject

import "time"

type AddFrRequest struct {
	RequestId    string    `json:"request_id"`
	AuthorId     string    `json:"author_id"`
	AccountId    string    `json:"account_id"` // Người được gửi
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`        // 3 loại: chờ / từ chối / chấp nhận
	ActiveStatus bool      `json:"active_status"` // Người gửi có thể hủy
}

func GetAddFriendRequestTable() string {
	return "add_friend_request"
}
