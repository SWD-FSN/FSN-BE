package service

import (
	"context"
	"database/sql"
	"log"
	business_object "social_network/business_object"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
	"time"
)

type notiService struct {
	userRepo repo.IUserRepo
	notiRepo repo.INotificationRepo
	logger   *log.Logger
}

func InitializeNotiService(db *sql.DB, logger *log.Logger) service.INotificationService {
	return &notiService{
		userRepo: repository.InitializeUserRepo(db, logger),
		notiRepo: repository.InitializeNotiRepo(db, logger),
		logger:   logger,
	}
}

func GenerateNotiService() (service.INotificationService, error) {
	var logger = util.GetLogConfig()

	db, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeNotiService(db, logger), nil
}

// CreateNotification implements service.INotificationService.
func (n *notiService) CreateNotification(req dto.CreateNotiRequest, ctx context.Context) error {
	return n.notiRepo.CreateNotification(business_object.Notification{
		NotificationId: util.GenerateId(),
		ActorId:        req.ActorId,
		ObjectId:       req.ObjectId,
		ObjectType:     req.ObjectType,
		Action:         req.Action,
		IsRead:         false,
		CreatedAt:      time.Now(),
	}, ctx)
}

// GetAllNotifications implements service.INotificationService.
func (n *notiService) GetAllNotifications(ctx context.Context) (*[]business_object.Notification, error) {
	res, err := n.notiRepo.GetAllNotifications(ctx)

	if err != nil {
		return nil, err
	}

	util.SortByTime(*res, func(noti business_object.Notification) time.Time { return noti.CreatedAt }, false)
	return res, nil
}

// GetUserNotifications implements service.INotificationService.
func (n *notiService) GetUserNotifications(id string, ctx context.Context) *dto.NotificationDialogResponseV2 {
	tmpStorage, _ := n.notiRepo.GetUserNotifications(id, ctx)

	var unreadNotis int = 0
	var notis []dto.NotificationResponseV2

	for _, noti := range *tmpStorage {
		// Get actor
		actor, _ := n.userRepo.GetUser(noti.ActorId, ctx)

		// Check actor
		if actor != nil {
			content, _ := generateContentAndContentTypeOfMsg(actor.Username, noti.Action, noti.ObjectType, "")
			notis = append(notis, dto.NotificationResponseV2{
				NotificationId: noti.NotificationId,
				Content:        content,
				CreatedAt:      noti.CreatedAt,
			})

			if !noti.IsRead {
				unreadNotis += 1
			}
		}
	}

	return &dto.NotificationDialogResponseV2{
		Notifications:    notis,
		UnreadNotiAmount: unreadNotis,
	}
}

// GetUserUnreadNotifications implements service.INotificationService.
func (n *notiService) GetUserUnreadNotifications(id string, ctx context.Context) (*[]business_object.Notification, error) {
	res, err := n.notiRepo.GetUserUnreadNotifications(id, ctx)

	if err != nil {
		return nil, err
	}

	util.SortByTime(*res, func(noti business_object.Notification) time.Time { return noti.CreatedAt }, false)
	return res, nil
}
