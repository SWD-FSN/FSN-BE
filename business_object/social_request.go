package businessobject

import "time"

// Delete after update request
type SocialRequest struct {
	RequestId   string    `json:"request_id"`
	AuthorId    string    `json:"author_id"`
	AccountId   string    `json:"account_id"`
	RequestType string    `json:"request_type"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetSocialRequestTable() string {
	return "social_request"
}
