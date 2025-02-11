package service

import (
	"context"
	business_object "social_network/business_object"
)

type IRoleService interface {
	GetAllRoles(ctx context.Context) (*[]business_object.Role, error)
	GetRolesByName(name string, ctx context.Context) (*[]business_object.Role, error)
	GetRolesByStatus(rawStatus string, ctx context.Context) (*[]business_object.Role, error)
	GetRoleById(id string, ctx context.Context) (*business_object.Role, error)
	CreateRole(name string, ctx context.Context) error
	UpdateRole(role business_object.Role, ctx context.Context) error
	RemoveRole(id string, ctx context.Context) error
	ActivateRole(id string, ctx context.Context) error
}
