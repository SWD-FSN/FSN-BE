package businessobject

import "time"

type Like struct {
	LikeId     string    `json:"like_id"`
	AuthorId   string    `json:"author_id"`
	ObjectId   string    `json:"object_id"`
	ObjectType string    `json:"object_type"` // Post, cmt, ...
	CreatedAt  time.Time `json:"created_at"`
}

func GetLikeTable() string {
	return "like"
}
