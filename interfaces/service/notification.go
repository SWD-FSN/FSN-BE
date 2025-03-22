package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type INotificationService interface {
	GetAllNotifications(ctx context.Context) (*[]business_object.Notification, error)
	GetUserNotifications(id string, ctx context.Context) *dto.NotificationDialogResponseV2
	GetUserUnreadNotifications(id string, ctx context.Context) (*[]business_object.Notification, error)
	CreateNotification(req dto.CreateNotiRequest, ctx context.Context) error
}
