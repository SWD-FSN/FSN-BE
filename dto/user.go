package dto

import "time"

type CreateUserReq struct {
	UserName      string    `json:"user_name"`
	RoleId        string    `json:"role_id"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	IsPrivate     *bool     `json:"is_private"`
	IsActive      *bool     `json:"is_active"`
}

type UserSaveModel struct {
	UserName      string    `json:"user_name"`
	RoleId        string    `json:"role_id"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	Followers     string    `jsonz:"followers"`
	Followings    string    `json:"followings"`
	BlockUsers    string    `json:"block_users"`
	Conversations string    `json:"conversations"`
	IsPrivate     bool      `json:"is_private"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserDBResModel struct {
	UserId        string    `json:"user_id"`
	RoleId        string    `json:"role_id"`
	UserName      string    `json:"user_name"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	Followers     string    `jsonz:"followers"`
	Followings    string    `json:"followings"`
	BlockUsers    string    `json:"block_users"`
	Conversations string    `json:"conversations"`
	IsPrivate     bool      `json:"is_private"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
