package repo

import (
	"context"
	business_object "social_network/business_object"
)

type ILikeRepo interface {
	GetAllLikes(ctx context.Context) (*[]business_object.Like, error)
	GetLikesFromObject(id, kind string, ctx context.Context) (*[]business_object.Like, error)
	GetLike(id string, ctx context.Context) (*business_object.Like, error)
	CreateLike(like business_object.Like, ctx context.Context) error
	CancelLike(id string, ctx context.Context) error
}
