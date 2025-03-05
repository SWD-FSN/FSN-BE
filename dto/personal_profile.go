package dto

type PersonalProfileUIResponse struct {
	UserId         string          `json:"user_id"`
	Username       string          `json:"username"`
	ProfileAvatar  string          `json:"profile_avatar"`
	IsFriend       bool            `json:"is_friend"`
	IsSelf         bool            `json:"is_self"`         // Tự vào xem acc của mình
	ConversationId string          `json:"conversation_id"` // Nút nhắn tin
	Posts          *[]PostResponse `json:"posts"`
}
