package dto

type UpPostReq struct {
	AuthorId  string `json:"author_id"`
	Content   string `json:"content"`
	IsPrivate *bool  `json:"is_private"`
	IsHidden  *bool  `json:"is_hidden"`
}

type UpdatePostReq struct {
	PostId    string `json:"post_id"`
	Content   string `json:"content" validate:"max=500"`
	IsPrivate *bool  `json:"is_private"`
	IsHidden  *bool  `json:"is_hidden"`
	Status    *bool  `json:"status"`
}
