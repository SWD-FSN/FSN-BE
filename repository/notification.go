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
	var query string

	var objectId string
	if notification.TargetUserId != "" {
		query = "INSERT INTO " + business_object.GetNotificationTable() + "(id, actor_id, target_user_id, action, is_read, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		objectId = notification.TargetUserId
	} else if notification.CommentId != "" {
		query = "INSERT INTO " + business_object.GetNotificationTable() + "(id, actor_id, comment_id, action, is_read, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		objectId = notification.CommentId
	} else if notification.PostId != "" {
		query = "INSERT INTO " + business_object.GetNotificationTable() + "(id, actor_id, post_id, action, is_read, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		objectId = notification.PostId
	}

	//defer n.db.Close()

	if _, err := n.db.Exec(query, notification.NotificationId, notification.ActorId, objectId, notification.Action, notification.IsRead, notification.CreatedAt); err != nil {
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

	//defer n.db.Close()

	rows, err := n.db.Query(query)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Notification
	for rows.Next() {
		var x business_object.Notification

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &x.TargetUserId, &x.PostId, &x.CommentId, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetNotificationOnAction implements repo.INotificationRepo.
func (n *notificationRepo) GetNotificationOnAction(req dto.GetNotiOnActionRequest, ctx context.Context) (*businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetNotificationOnAction - "
	var query string = "SELECT * FROM " + business_object.GetNotificationTable() + " WHERE actor_id = $1 AND created_at = $2 LIMIT 1"

	//defer n.db.Close()

	var res business_object.Notification
	var targetUserId, postId, commentId sql.NullString

	if err := n.db.QueryRow(query, req.ActorId, req.CreatedAt).Scan(&res.NotificationId, &res.ActorId, &targetUserId, &postId, &commentId, &res.Action, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		n.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	if targetUserId.Valid {
		res.TargetUserId = targetUserId.String
	}

	if postId.Valid {
		res.PostId = postId.String
	}

	if commentId.Valid {
		res.CommentId = commentId.String
	}

	return &res, nil
}

// GetNotification implements repo.INotificationRepo.
func (n *notificationRepo) GetNotification(id string, ctx context.Context) (*businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetNotification - "
	var query string = "SELECT * FROM " + business_object.GetNotificationTable() + " WHERE id = $1"

	//defer n.db.Close()

	var res business_object.Notification
	var targetUserId, postId, commentId sql.NullString

	if err := n.db.QueryRow(query, id).Scan(&res.NotificationId, &res.ActorId, &targetUserId, &postId, &commentId, &res.Action, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		n.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	if targetUserId.Valid {
		res.TargetUserId = targetUserId.String
	}

	if postId.Valid {
		res.PostId = postId.String
	}

	if commentId.Valid {
		res.CommentId = commentId.String
	}

	return &res, nil
}

// GetUserNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetUserNotifications(id string, ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetUserNotifications - "
	var query string = `
		SELECT n.id, n.actor_id, n.target_user_id, 
       			n.post_id, n.comment_id, n.action, 
       			n.is_read, n.created_at
		FROM notifications n
		JOIN posts p ON n.post_id = p.id
		WHERE p.author_id = $1

		UNION 

		SELECT n.id, n.actor_id, n.target_user_id, 
       			n.post_id, n.comment_id, n.action, 
       			n.is_read, n.created_at
		FROM notifications n
		JOIN comments c ON n.comment_id = c.id
		WHERE c.author_id = $2

		UNION

		SELECT n.id, n.actor_id, n.target_user_id, 
       			n.post_id, n.comment_id, n.action, 
       			n.is_read, n.created_at
		FROM notifications n
		WHERE n.target_user_id = $3

		ORDER BY created_at DESC;
	`

	var internalErr error = errors.New(noti.InternalErr)
	//defer n.db.Close()

	rows, err := n.db.Query(query, id, id, id)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Notification
	for rows.Next() {
		var x business_object.Notification
		var targetUserId, postId, commentId sql.NullString

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &targetUserId, &postId, &commentId, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		if targetUserId.Valid {
			x.TargetUserId = targetUserId.String
		}

		if postId.Valid {
			x.PostId = postId.String
		}

		if commentId.Valid {
			x.CommentId = commentId.String
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetUserUnreadNotifications implements repo.INotificationRepo.
func (n *notificationRepo) GetUserUnreadNotifications(id string, ctx context.Context) (*[]businessobject.Notification, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "GetUserNotifications - "
	var query string = `
		SELECT n.*
		FROM notifications n
		JOIN posts p ON n.post_id = p.id
		WHERE n.is_read = false AND p.author_id = $1

		UNION 
		
		SELECT n.*
		FROM notifications n
		JOIN comments c ON n.comment_id = c.id
		WHERE n.is_read = false AND c.author_id = $2

		UNION
		SELECT n.*
		FROM notifications n
		JOIN users u on n.target_user_id = u.id
		WHERE n.is_read = false AND u.id = $3

		ORDER BY n.created_at DESC
	`
	var internalErr error = errors.New(noti.InternalErr)

	//defer n.db.Close()

	rows, err := n.db.Query(query, id, id, id)
	if err != nil {
		n.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Notification
	for rows.Next() {
		var x business_object.Notification
		var targetUserId, postId, commentId sql.NullString

		if err := rows.Scan(&x.NotificationId, &x.ActorId, &targetUserId, &postId, &commentId, &x.Action, &x.IsRead, &x.CreatedAt); err != nil {
			n.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		if targetUserId.Valid {
			x.TargetUserId = targetUserId.String
		}

		if postId.Valid {
			x.PostId = postId.String
		}

		if commentId.Valid {
			x.CommentId = commentId.String
		}

		res = append(res, x)
	}

	return &res, nil
}

// RemoveNotification implements repo.INotificationRepo.
func (n *notificationRepo) RemoveNotification(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetNotificationTable()) + "RemoveNotification - "
	var query string = "DELETE FROM " + business_object.GetNotificationTable() + " WHERE id = $1"
	var internalErrMsg error = errors.New(noti.InternalErr)

	//defer n.db.Close()

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

	//defer n.db.Close()

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
