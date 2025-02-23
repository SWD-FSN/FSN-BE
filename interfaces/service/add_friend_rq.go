package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type IAddFriendRqService interface {
	GetAllAddFrRqs(ctx context.Context) (*[]business_object.Follow, error)
	GetUserAddFrRqs(id string, ctx context.Context) (*[]business_object.Follow, error)   // Xem request của user đã gửi
	GetAddFrRqsToUser(id string, ctx context.Context) (*[]business_object.Follow, error) // Xem request đc gửi đến 1 user - Bản thân xem có những ai đã gửi lời mời kết bạn đến mình
	GetAddFrRq(id string, ctx context.Context) (*business_object.Follow, error)
	ProcessAddFrRq(req dto.SocialRequest, ctx context.Context) error
	AcceptAddFrRq(requestId, actorId string, ctx context.Context) error
	CancelAddFrRq(requestId, actorId string, ctx context.Context) error
}
