package repo

import (
	"context"
	business_object "social_network/business_object"
)

type IPostRepo interface {
	GetAllPosts(ctx context.Context) (*[]business_object.Post, error)
	GetPostsByUser(id string, ctx context.Context) (*[]business_object.Post, error)
	GetPostsByKeyword(keyword string, ctx context.Context) (*[]business_object.Post, error)
	GetPost(id string, ctx context.Context) (*business_object.Post, error)
	CreatePost(post business_object.Post, ctx context.Context) error
	UpdatePost(post business_object.Post, ctx context.Context) error
	RemovePost(id string, ctx context.Context) error
}
