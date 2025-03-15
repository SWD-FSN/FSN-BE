package dto

type DoLikeReq struct {
	ActorId    string `json:"actor_id"`
	ObjectId   string `json:"object_id"`
	ObjectType string `json:"object_type"`
}
