package businessobject

import "time"

type AddFrRequest struct {
	RequestId string    `json:"request_id"`
	AuthorId  string    `json:"author_id"`
	AccountId string    `json:"account_id"` // Người được gửi
	CreatedAt time.Time `json:"created_at"`
}

func GetAddFriendRequestTable() string {
	return "add_friend_request"
}
