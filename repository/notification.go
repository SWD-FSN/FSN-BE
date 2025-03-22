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
	"social_network/dto"
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
	var query string = "INSERT INTO " + business_object.GetNotificationTable() + "(id, actor_id, object_id, object_type, action, is_read, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	defer n.db.Close()

	if _, err := n.db.Exec(query, notification.NotificationId, notification.ActorId, notification.ObjectId, notification.ObjectType, notification.Action, notification.IsRead, notification.CreatedAt); err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAllNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetAllNotifications(ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetAllNotifications - "
	var query string = "SELECT * FROM " + business_object.GetNotificationTable()
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

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &x.ObjectId, &x.ObjectType, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetNotificationOnAction implements repo.INotificationRepo.
func (n *notificationRepo) GetNotificationOnAction(req dto.GetNotiOnActionRequest, ctx context.Context) (*businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetNotificationOnAction - "
	var query string = "SELECT * FROM " + business_object.GetNotificationTable() + " WHERE actor_id = $1, object_id = $2, object_type = $3, action = $4 and created_at = $5"

	defer n.db.Close()

	var res *business_object.Notification
	if err := n.db.QueryRow(query, req.ActorId, req.ObjectId, req.ObjectType, req.Action, req.CreatedAt).Scan(&res.NotificationId, &res.ActorId, &res.ObjectId, &res.ObjectType, &res.Action, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		n.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetNotification implements repo.INotificationRepo.
func (n *notificationRepo) GetNotification(id string, ctx context.Context) (*businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetNotification - "
	var query string = "SELECT * FROM " + business_object.GetNotificationTable() + " WHERE id = $1"

	defer n.db.Close()

	var res *business_object.Notification
	if err := n.db.QueryRow(query, id).Scan(&res.NotificationId, &res.ActorId, &res.ObjectId, &res.ObjectType, &res.Action, &res.CreatedAt); err != nil {
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
	var query string = `
		SELECT n.*
		FROM notifications n
		JOIN posts p ON n.object_id = p.id
		WHERE p.author_id = $1

		UNION 
		
		SELECT n.*
		FROM notifications n
		JOIN comments c ON n.object_id = c.id
		WHERE c.author_id = $2

		ORDER BY n.created_at DESC
	`

	var internalErr error = errors.New(noti.InternalErr)
	defer n.db.Close()

	rows, err := n.db.Query(query, id, id)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Notification
	for rows.Next() {
		var x business_object.Notification

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &x.ObjectId, &x.ObjectType, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetUserUnreadNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetUserUnreadNotifications(id string, ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetUserNotifications - "
	var query string = `
		SELECT n.*
		FROM notifications n
		JOIN posts p ON n.object_id = p.id
		WHERE n.is_read = false AND p.author_id = $1

		UNION 
		
		SELECT n.*
		FROM notifications n
		JOIN comments c ON n.object_id = c.id
		WHERE n.is_read = false AND c.author_id = $2

		ORDER BY n.created_at DESC
	`
	var internalErr error = errors.New(noti.InternalErr)

	defer n.db.Close()

	rows, err := n.db.Query(query, id, id)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Notification
	for rows.Next() {
		var x business_object.Notification

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &x.ObjectId, &x.ObjectType, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// RemoveNotification implements repo.INotificationRepo.
func (n *notificationRepo) RemoveNotification(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "RemoveNotification - "
	var query string = "DELETE FROM " + business_object.GetNotificationTable() + " WHERE id = $1"
	var internalErrMsg error = errors.New(noti.InternalErr)

	defer n.db.Close()

	res, err := n.db.Exec(query, id)

	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetNotificationTable()))
	}

	return nil
}

// NoteReadNotification implements repo.INotificationRepo.
func (n *notificationRepo) NoteReadNotification(id string, ctx context.Context) error {
	var query string = "UPDATE " + business_object.GetNotificationTable() + " SET is_read = true WHERE id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "NoteReadNotification - "
	var internalErrMsg error = errors.New(noti.InternalErr)

	defer n.db.Close()

	res, err := n.db.Exec(query, id)

	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetNotificationTable()))
	}

	return nil
}

// UpdateNotification implements repo.INotificationRepo. -- Hiện tại chưa cần xài đến
func (n *notificationRepo) UpdateNotification(notification businessobject.Notification, ctx context.Context) error {
	panic("unimplemented")
}
