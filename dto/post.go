package dto

import "time"

type UpPostReq struct {
	AuthorId   string `json:"author_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	Attachment string `json:"attachment"`
	IsPrivate  *bool  `json:"is_private"`
	IsHidden   *bool  `json:"is_hidden"`
}

type UpdatePostReq struct {
	PostId     string `json:"post_id"`
	Content    string `json:"content" validate:"max=500"`
	Attachment string `json:"attachment"`
	IsPrivate  *bool  `json:"is_private"`
	IsHidden   *bool  `json:"is_hidden"`
	Status     *bool  `json:"status"`
}

type PostResponse struct {
	// Primary Post Properties
	PostId     string    `json:"post_id"`
	Content    string    `json:"content" validate:"max=500"`
	Attachment string    `json:"attachment"`
	IsPrivate  bool      `json:"is_private"`
	IsHidden   bool      `json:"is_hidden"`
	LikeAmount int       `json:"like_amount"`
	CreatedAt  time.Time `json:"created_at"`
	// User Properties
	AuthorId      string `json:"author_id"`
	Username      string `json:"username"`
	ProfileAvatar string `json:"profile_avatar"`
}
