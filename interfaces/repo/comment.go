package repo

import (
	"context"
	business_object "social_network/business_object"
)

type ICommentRepo interface {
	CreateComment(cmt business_object.Comment, ctx context.Context) error
	GetComment(id string, ctx context.Context) (*business_object.Comment, error)
	GetCommentsFromPost(id string, ctx context.Context) (*[]business_object.Comment, error)
	RemoveComment(id string, ctx context.Context) error
	EditComment(cmt business_object.Comment, ctx context.Context) error
}
