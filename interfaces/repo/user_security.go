package repo

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type IUserSecurityRepo interface {
	GetUserSecurity(id string, ctx context.Context) (*business_object.UserSecurity, error)
	CreateUserSecurity(usc business_object.UserSecurity, ctx context.Context) error
	EditUserSecurity(usc business_object.UserSecurity, ctx context.Context) error
	Login(req dto.LoginSecurityRequest, ctx context.Context) error
	LogOut(id string, ctx context.Context) error
}
