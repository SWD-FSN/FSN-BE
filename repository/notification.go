package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	business_object "social_network/business_object"
	businessobject "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
)

type notificationRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeNotiRepo(db *sql.DB, logger *log.Logger) repo.INotificationRepo {
	return &notificationRepo{
		db:     db,
		logger: logger,
	}
}

// CreateNotification implements repo.INotificationRepo.
func (n *notificationRepo) CreateNotification(notification businessobject.Notification, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "CreateNotification - "
	var query string = "Insert into " + business_object.GetSocialRequestTable() + "(id, object_id, object_type, action, is_read, created_at) values (?, ?, ?, ?, ?, ?)"

	defer n.db.Close()

	if _, err := n.db.Exec(query, notification.NotificationId, notification.ObjectId, notification.ObjectType, notification.Action, notification.IsRead, notification.CreatedAt); err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil

}

// GetAllNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetAllNotifications(ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetAllNotifications - "
	var query string = "Select * from " + business_object.GetNotificationTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer n.db.Close()

	rows, err := n.db.Query(query)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Notification
	for rows.Next() {
		var x business_object.Notification

		if err := rows.Scan(&x.NotificationId, &x.ObjectId, &x.ObjectType, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetNotification implements repo.INotificationRepo.
func (n *notificationRepo) GetNotification(id string, ctx context.Context) (*businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetNotification - "
	var query string = "Select * from " + business_object.GetNotificationTable() + " where id = ?"

	defer n.db.Close()

	var res *business_object.Notification
	if err := n.db.QueryRow(query, id).Scan(&res.NotificationId, &res.ObjectId, &res.ObjectType, &res.Action, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		n.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetUserNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetUserNotifications(id string, ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetUserNotifications - "
	var query string = "Select * from " + business_object.GetNotificationTable() + " where "
	var internalErr error = errors.New(noti.InternalErr)

	defer n.db.Close()

	rows, err := n.db.Query(query)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Notification
	for rows.Next() {
		var x business_object.Notification

		if err := rows.Scan(&x.NotificationId, &x.ObjectId, &x.ObjectType, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// RemoveNotification implements repo.INotificationRepo.
func (n *notificationRepo) RemoveNotification(id string, ctx context.Context) error {
	panic("unimplemented")
}

// UpdateNotification implements repo.INotificationRepo.
func (n *notificationRepo) UpdateNotification(notification businessobject.Notification, ctx context.Context) error {
	panic("unimplemented")
}
