package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	businessObject "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
	"time"
)

type conversationRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeConversationRepo(db *sql.DB, logger *log.Logger) repo.IConversationRepo {
	return &conversationRepo{
		db:     db,
		logger: logger,
	}
}

// GetConversationOfTwoUsers implements repo.IConversationRepo.
func (c *conversationRepo) GetConversationOfTwoUsers(userId1, userId2 string, ctx context.Context) (*businessObject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "GetConversationOfTwoUsers - "
	var query string = "SELECT * FROM " + businessObject.GetConversationTable() + " WHERE members LIKE '%' || $1 || '%' AND members LIKE '%' || $2 || '%'"
	var internalErr error = errors.New(noti.InternalErr)

	rows, err := c.db.Query(query, userId1, userId2)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	for rows.Next() {
		var x businessObject.Conversation

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &x.IsDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		// Not a group chat
		if x.HostId == nil || !x.IsGroup {
			return &x, nil
		}
	}

	return nil, nil
}

// UpdateConversation implements repo.IConversationRepo.
func (c *conversationRepo) UpdateConversation(conversation businessObject.Conversation, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "UpdateConversation - "
	var query string = "UPDATE " + businessObject.GetConversationTable() + " SET conversation_avatar = ?, name = ?, host_id = ?, members = ?, is_delete = ? AND updated_at = ? WHERE id = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	res, err := c.db.Exec(query, conversation.ConversationAvatar, conversation.ConversationName, conversation.HostId, conversation.Members, conversation.IsDelete, conversation.UpdatedAt, conversation.ConversationId)
	if err != nil {
		c.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, businessObject.GetConversationTable()))
	}

	return nil
}

// CreateConversation implements repo.IConversationRepo.
func (c *conversationRepo) CreateConversation(conversation businessObject.Conversation, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "CreateConversation - "
	var query string = "INSERT INTO " + businessObject.GetConversationTable() + "(id, name, host_id, members, is_group, is_delete, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	defer c.db.Close()

	if _, err := c.db.Exec(query, conversation.ConversationId, conversation.ConversationName, conversation.HostId, conversation.Members, conversation.IsGroup, conversation.IsDelete, conversation.CreatedAt, conversation.UpdatedAt); err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// DissolveGroupConversation implements repo.IConversationRepo.
func (c *conversationRepo) DissolveGroupConversation(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "DissovelGroupConversation - "
	var query string = "UPDATE " + businessObject.GetConversationTable() + " SET status = false, updated_at = $1 WHERE id = $2"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	res, err := c.db.Exec(query, time.Now().String(), id)
	if err != nil {
		c.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, businessObject.GetConversationTable()))
	}

	return nil
}

// GetAllConversations implements repo.IConversationRepo.
func (c *conversationRepo) GetAllConversations(ctx context.Context) (*[]businessObject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "GetAllConversations - "
	var query string = "SELECT * FROM " + businessObject.GetConversationTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []businessObject.Conversation
	for rows.Next() {
		var x businessObject.Conversation
		var isDelete sql.NullBool

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &isDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		if isDelete.Valid {
			*x.IsDelete = isDelete.Bool
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetConversation implements repo.IConversationRepo.
func (c *conversationRepo) GetConversation(id string, ctx context.Context) (*businessObject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "GetConversation - "
	var query string = "SELECT * FROM " + businessObject.GetConversationTable() + " WHERE id = $1"

	defer c.db.Close()

	var res businessObject.Conversation
	var isDelete sql.NullBool
	if err := c.db.QueryRow(query, id).Scan(&res.ConversationId, &res.ConversationName, &res.HostId, &res.Members, &res.IsGroup, &isDelete, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		c.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	if isDelete.Valid {
		*res.IsDelete = isDelete.Bool
	}

	return &res, nil
}

// GetConversationsByKeyword implements repo.IConversationRepo.
func (c *conversationRepo) GetConversationsByKeyword(id string, keyword string, ctx context.Context) (*[]businessObject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "GetConversationsFROMUser - "
	var query string = "SELECT * FROM " + businessObject.GetConversationTable() + " WHERE members like '%$1%' AND lower(name) like lower('%$2%')"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query, id, keyword)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []businessObject.Conversation
	for rows.Next() {
		var x businessObject.Conversation
		var isDelete sql.NullBool

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &isDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		if isDelete.Valid {
			*x.IsDelete = isDelete.Bool
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetConversationsFROMUser implements repo.IConversationRepo.
func (c *conversationRepo) GetConversationsFromUser(id string, ctx context.Context) (*[]businessObject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, businessObject.GetConversationTable()) + "GetConversationsFROMUser - "
	var query string = "SELECT * FROM " + businessObject.GetConversationTable() + " WHERE members LIKE '%$1%'"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query, id)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []businessObject.Conversation
	for rows.Next() {
		var x businessObject.Conversation
		var isDelete sql.NullBool
		var hostId sql.NullString

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &hostId, &x.Members, &x.IsGroup, &isDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		if isDelete.Valid {
			*x.IsDelete = isDelete.Bool
		}

		if hostId.Valid {
			x.HostId = &hostId.String
		}

		res = append(res, x)
	}

	return &res, nil
}
