package dto

import (
	"net/http"
	"time"
)

type CreateUserReq struct {
	Username      string    `json:"username" validate:"required"`
	RoleId        string    `json:"role_id"`
	FullName      string    `json:"full_name" validate:"required"`
	Email         string    `json:"email" validate:"required"`
	Password      string    `json:"password" validate:"required,min=8"`
	DateOfBirth   time.Time `json:"date_of_birth" validate:"required"`
	ProfileAvatar string    `json:"profile_avatar"`
}

type UpdateUserReq struct {
	UserId        string     `json:"user_id"`
	RoleId        string     `json:"role_id"`
	FullName      string     `json:"full_name"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	Password      string     `json:"password" validate:"min=10"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	ProfileAvatar string     `json:"profile_avatar"`
	Bio           string     `json:"bio"`
	IsPrivate     *bool      `json:"is_private"`
}

type UserDBResModel struct {
	UserId          string    `json:"user_id"`
	RoleId          string    `json:"role_id"`
	FullName        string    `json:"full_name"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        string    `json:"password" validate:"min=10"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	ProfileAvatar   string    `json:"profile_avatar"`
	Bio             string    `json:"bio"`
	Friends         string    `json:"friends"`
	Followers       string    `json:"followers"`
	Followings      string    `json:"followings"`
	BlockUsers      string    `json:"block_users"`
	Conversations   string    `json:"conversations"`
	IsPrivate       bool      `json:"is_private"`
	IsActive        bool      `json:"is_active"`
	IsActivated     bool      `json:"is_activated"`
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

type GetInvoledAccouuntsRequest struct {
	UserId       string `json:"user_id"`
	InvolvedType string `json:"involed_type"`
}

type GetInvolvedAccountsSearchResponse struct {
	UserId        string `json:"user_id"`
	Username      string `json:"username"`
	ProfileAvatar string `json:"profile_avatar"`
}

type UserSearchDoneResponse struct {
	UserId            string `json:"user_id"`
	Username          string `json:"username"`
	ProfileAvatar     string `json:"profile_avatar"`
	FollowerAmount    int    `json:"follower_amount"`
	IsFriendWithActor bool   `json:"is_friend_with_actor"`
}

type UserConnectionRequest struct {
	UserId  string
	Request *http.Request
	Writer  http.ResponseWriter
}
