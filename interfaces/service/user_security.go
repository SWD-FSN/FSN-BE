package service

import (
	"context"
	business_object "social_network/business_object"
)

type IUserSecurityService interface {
	GetUserSecurity(id string, ctx context.Context) (*business_object.UserSecurity, error)
	EditUserSecurity(usc business_object.UserSecurity, ctx context.Context) error
}
