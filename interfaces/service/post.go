package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type IPostService interface {
	GetAllPosts(ctx context.Context) (*[]business_object.Post, error)
	GetPostsByUser(id string, ctx context.Context) (*[]business_object.Post, error)
	GetPosts(ctx context.Context) *[]dto.PostResponse
	GetPost(id string, ctx context.Context) (*business_object.Post, error)
	UpPost(req dto.UpPostReq, ctx context.Context) error
	UpdatePost(req dto.UpdatePostReq, ctx context.Context) error
	RemovePost(id string, actorId string, ctx context.Context) error
}
