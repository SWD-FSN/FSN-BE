package businessobject

import "time"

type Comment struct {
	CommentId string    `json:"comment_id"`
	AuthorId  string    `json:"author_id"`
	PostId    string    `json:"post_id"`
	Content   string    `json:"content" validate:"max=150"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetCommentTable() string {
	return "comments"
}
