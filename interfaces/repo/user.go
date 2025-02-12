package repo

import (
	"context"
	"social_network/dto"
)

type IUserRepo interface {
	GetAllUsers(ctx context.Context) (*[]dto.UserDBResModel, error)
	GetUsersByRole(id string, ctx context.Context) (*[]dto.UserDBResModel, error)
	GetUsersByStatus(status bool, ctx context.Context) (*[]dto.UserDBResModel, error)
	GetUser(id string, ctx context.Context) (*dto.UserDBResModel, error)
	GetUserByEmail(email string, ctx context.Context) (*dto.UserDBResModel, error)
	CreateUser(user dto.UserSaveModel, ctx context.Context) error
	UpdateUser(user dto.UserDBResModel) error
	ChangeUserStatus(id string, status bool, ctx context.Context) error
}
