package businessobject

type Like struct {
	LikeId   string `json:"like_id"`
	AuthorId string `json:"author_id"`
	ObjectId string `json:"object_id"`
	Status   bool   `json:"status"`
}
