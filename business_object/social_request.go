package businessobject

import "time"

// Delete after update request
type SocialRequest struct { // Request kết bạn, request foilow
	RequestId   string    `json:"request_id"`
	AuthorId    string    `json:"author_id"`    // Người gửi request
	AccountId   string    `json:"account_id"`   // Người nhận request
	RequestType string    `json:"request_type"` // KẾt bạn hoặc follow
	CreatedAt   time.Time `json:"created_at"`
}

func GetSocialRequestTable() string {
	return "social_request"
}
