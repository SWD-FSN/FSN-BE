package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type IUserService interface {
	GetUser(id string, ctx context.Context) (*business_object.User, error)
	CreateUser(req dto.CreateUserReq, ctx context.Context) error
}
