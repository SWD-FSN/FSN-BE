package dto

import "time"

type CreateUserReq struct {
	Username      string    `json:"username"`
	RoleId        string    `json:"role_id"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	IsPrivate     *bool     `json:"is_private"`
	IsActive      *bool     `json:"is_active"`
}

type UserSaveModel struct {
	Username      string    `json:"username"`
	RoleId        string    `json:"role_id"`
	Email         string    `json:"email"`
	Password      string    `json:"password" validate:"min=10"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	Friends       string    `json:"friends"`
	Followers     string    `jsonz:"followers"`
	Followings    string    `json:"followings"`
	BlockUsers    string    `json:"block_users"`
	Conversations string    `json:"conversations"`
	IsPrivate     bool      `json:"is_private"`
	IsActive      bool      `json:"is_active"`
	IsActivated   bool      `json:"json:is_activated"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserDBResModel struct {
	UserId          string    `json:"user_id"`
	RoleId          string    `json:"role_id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        string    `json:"password" validate:"min=10"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	ProfileAvatar   string    `json:"profile_avatar"`
	Bio             string    `json:"bio"`
	Friends         string    `json:"friends"`
	Followers       string    `jsonz:"followers"`
	Followings      string    `json:"followings"`
	BlockUsers      string    `json:"block_users"`
	Conversations   string    `json:"conversations"`
	IsPrivate       bool      `json:"is_private"`
	IsActive        bool      `json:"is_active"`
	IsActivated     bool      `json:"json:is_activated"`
	IsHaveToResetPw *bool     `json:"is_have_to_reset_password"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginSecurityRequest struct {
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type VerifyAccountRequest struct {
}
