package businessobject

import "time"

type Comment struct {
	CommentId       string        `json:"comment_id"`
	ParentCommentId string        `json:"parent_comment_id"`
	AuthorId        string        `json:"author_id"`
	PostId          string        `json:"post_id"`
	Attachments     *[]Attachment `json:"attachments"`
	Content         string        `json:"content" validate:"max=150"`
	Tags            *[]string     `json:"tags"`
	Mentions        *[]string     `json:"mentions"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	IsEdited        bool          `json:"is_edited"`
	Status          bool          `json:"status"`
}

func GetCommentTable() string {
	return "comment"
}
