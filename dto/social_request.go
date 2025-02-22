package dto

type ActionRequest struct {
	AuthorId   string `json:"author_id"`
	AccountId  string `json:"account_id"`
	ActionType string `json:"action_type"` // follow, like
}
