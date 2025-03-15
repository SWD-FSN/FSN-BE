package dto

import "time"

// Request
type CreateCommentRequest struct {
	ActorId string `json:"actor_id"`
	PostId  string `json:"post_id"`
	Content string `json:"content" validate:"required"`
}

type EditCommentRequest struct {
	ActorId   string `json:"actor_id"`
	CommentId string `json:"comment_id"`
	Content   string `json:"content" validate:"required"`
}

// Response
type CommentDataResponse struct {
	AuthorId      string    `json:"author_id"`
	Author_avatar string    `json:"author_avatar"`
	Username      string    `json:"username"`
	CommentId     string    `json:"comment_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	LikeAmount    int       `json:"like_amount"`
}
