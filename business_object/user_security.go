package businessobject

import "time"

type UserSecurity struct {
	UserId            string    `json:"user_id"`
	AccessToken       string    `json:"access_token"`
	AccessEpiration   time.Time `json:"access_expiration"`
	RefreshToken      string    `json:"refresh_token"`
	RefreshExpiration time.Time `json:"refresh_expiration"`
	ActionToken       string    `json:"action_token"` // Lưu trữ token cho thay đổi mail, reset pass
	ActionExpiration  string    `json:"action_Expiration"`
	FailAccess        int       `json:"fail_access"`
	LastFail          time.Time `json:"last_fail"`
}

func GetUserSecurityTable() string {
	return "user_security"
}
