package businessobject

import "time"

// Delete after update request
type Follow struct {
	FollowId          string    `json:"post_id"`
	AuthorId          string    `json:"author_id"`
	AffectedAccountId string    `json:"affected_account_id"`
	CreatedAt         time.Time `json:"create_at"`
}

func GetFollowTable() string { return "follow" }
