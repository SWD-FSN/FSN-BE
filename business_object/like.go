package businessobject

import "time"

type Like struct {
	LikeId    string    `json:"like_id"`
	AuthorId  string    `json:"author_id"` // NGười thực hiện hành động like
	PostId    string    `json:"post_id"`
	CommentId string    `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
}

func GetLikeTable() string {
	return "likes"
}
