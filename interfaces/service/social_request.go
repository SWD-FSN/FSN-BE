package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type ISocialRequestService interface {
	GetAllRequests(ctx context.Context) (*[]business_object.SocialRequest, error)
	GetUserRequests(id, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error)   // Xem request của user đã gửi
	GetRequestsToUser(id, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error) // Xem request đc gửi đến 1 user - Bản thân xem có những ai đã gửi lời mời kết bạn đến mình
	GetRequest(id string, ctx context.Context) (*business_object.SocialRequest, error)
	ProcessRequest(req dto.SocialRequest, ctx context.Context) error
	AcceptRequest(requestId, actorId string, ctx context.Context) error
	CancelRequest(requestId, actorId string, ctx context.Context) error
}
