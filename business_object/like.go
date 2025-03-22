package businessobject

import "time"

type Like struct {
	LikeId     string    `json:"like_id"`
	AuthorId   string    `json:"author_id"`   // NGười thực hiện hành động like
	ObjectId   string    `json:"object_id"`   // Bài post, 1 comment
	ObjectType string    `json:"object_type"` // Post, cmt, ...
	CreatedAt  time.Time `json:"created_at"`
}

func GetLikeTable() string {
	return "likes"
}
