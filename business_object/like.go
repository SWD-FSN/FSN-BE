package businessobject

type Like struct {
	LikeId    string `json:"like_id"`
	AuthorId  string `json:"author_id"`
	AccountId string `json:"account_id"` // Tài khoản đc tác động đến
	ObjectId  string `json:"object_id"`
	Status    bool   `json:"status"`
}

func GetLikeTable() string {
	return "like"
}
