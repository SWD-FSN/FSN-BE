package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type ILikeService interface {
	GetAllLikes(ctx context.Context) (*[]business_object.Like, error)
	GetLikesFromObject(id, kind string, ctx context.Context) (*[]business_object.Like, error)
	GetLike(id string, ctx context.Context) (*business_object.Like, error)
	DoLike(req dto.DoLikeReq, ctx context.Context) error
	UndoLike(id string, ctx context.Context) error
}
