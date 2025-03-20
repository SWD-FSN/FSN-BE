package businessobject

import "time"

type User struct {
	UserId          string    `json:"user_id"`
	Username        string    `json:"username"`
	RoleId          string    `json:"role_id"`
	FullName        string    `json:"full_name"`
	Email           string    `json:"email" validate:"email, required"`
	Password        string    `json:"password" validate:"min=10"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	ProfileAvatar   string    `json:"profile_avatar"`
	Bio             string    `json:"bio"`
	Friends         *[]string `json:"friends"`
	Followers       *[]string `json:"followers"`
	Followings      *[]string `json:"followings"`
	BlockUsers      *[]string `json:"block_users"`
	IsPrivate       *bool     `json:"is_private"`
	IsActive        *bool     `json:"is_active"`
	IsActivated     bool      `json:"json:is_activated"`
	IsHaveToResetPw *bool     `json:"is_have_to_reset_password"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func GetUserTable() string {
	return "user"
}
