package repo

import (
	"context"
	business_object "social_network/business_object"
)

type INotificationRepo interface {
	GetAllNotifications(ctx context.Context) (*[]business_object.Notification, error)
	GetUserNotifications(id string, ctx context.Context) (*[]business_object.Notification, error)
	GetNotification(id string, ctx context.Context) (*business_object.Notification, error)
	CreateNotification(notification business_object.Notification, ctx context.Context) error
	UpdateNotification(notification business_object.Notification, ctx context.Context) error
	RemoveNotification(id string, ctx context.Context) error
}
