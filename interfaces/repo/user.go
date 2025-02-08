package repo

import (
	"context"
	"social_network/dto"
)

type IUserRepo interface {
	GetUser(id string, ctx context.Context) (*dto.UserDBResModel, error)
	CreateUser(user dto.UserSaveModel, ctx context.Context) error
}
