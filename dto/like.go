package dto

type DoLikeReq struct {
	AuthorId   string `json:"author_id"`
	ObjectId   string `json:"object_id"`
	ObjectType string `json:"object_type"`
}
