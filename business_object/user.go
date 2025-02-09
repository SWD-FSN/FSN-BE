package businessobject

import "time"

type User struct {
	UserId        string    `json:"user_id"`
	RoleId        string    `json:"role_id"`
	Username      string    `json:"username"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ProfileAvatar string    `json:"profile_avatar"`
	Bio           string    `json:"bio"`
	Friends       *[]string `json:"friends"`
	Followers     *[]string `json:"followers"`
	Followings    *[]string `json:"followings"`
	BlockUsers    *[]string `json:"block_users"`
	IsPrivate     *bool     `json:"is_private"`
	IsActive      *bool     `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetUserTable() string {
	return "user"
}
