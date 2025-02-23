package repo

import (
	"context"
	business_object "social_network/business_object"
)

type IAddFriendRqRepo interface {
	GetAllAddFrRqs(ctx context.Context) (*[]business_object.Follow, error)
	GetUserAddFrRqs(id string, ctx context.Context) (*[]business_object.Follow, error)   // Xem request của user đã gửi
	GetAddFrRqsToUser(id string, ctx context.Context) (*[]business_object.Follow, error) // Xem request đc gửi đến 1 user - Bản thân xem có những ai đã gửi lời mời kết bạn đến mình
	GetAddFrRq(id string, ctx context.Context) (*business_object.Follow, error)
	CreateAddFrRq(req business_object.Follow, ctx context.Context) error
	RemoveAddFrRq(id string, ctx context.Context) error
}
