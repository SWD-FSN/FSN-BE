package businessobject

import "time"

// Delete after update request
type SocialRequest struct {
	FollowId    string    `json:"post_id"`
	AuthorId    string    `json:"author_id"`
	AccountId   string    `json:"account_id"`
	RequestType string    `json:"request_type"`
	CreatedAt   time.Time `json:"create_at"`
}

func GetFollowTable() string { return "SocialRequest" }
