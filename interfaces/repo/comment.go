package repo

import (
	"context"
	business_object "social_network/business_object"
)

type ICommentRepo interface {
	GetComment(id string, ctx context.Context) (*business_object.Comment, error)
}
