package service

import (
	"context"
	"social_network/dto"
)

type ICommentService interface {
	PostComment(req dto.CreateCommentRequest, ctx context.Context) error
	GetCommentsFromPost(id string, ctx context.Context) *[]dto.CommentDataResponse
	RemoveComment(actorId, commentId string, ctx context.Context) error
	EditComment(req dto.EditCommentRequest, ctx context.Context) error
}
