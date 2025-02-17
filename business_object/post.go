package businessobject

import "time"

type Post struct {
	PostId    string    `json:"post_id"`
	AuthorId  string    `json:"author_id"`
	Content   string    `json:"content" validate:"max=500"`
	IsPrivate bool      `json:"is_private"`
	IsHidden  bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status"`
}

func GetPostTable() string {
	return "post"
}
