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
	"social_network/util"
	"strings"
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

// CreateConversation implements repo.IConversationRepo.
func (c *conversationRepo) CreateConversation(conversation businessobject.Conversation, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "CreateConversation - "
	var query string = "Insert into " + business_object.GetConversationTable() + "(id, conversation_name, host_id, members, is_group, is_delete, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)"

	defer c.db.Close()

	if _, err := c.db.Exec(query, conversation.ConversationId, conversation.ConversationName, conversation.HostId, conversation.Members, conversation.IsGroup, conversation.IsDelete, conversation.CreatedAt, conversation.UpdatedAt); err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// DissovelGroupConversation implements repo.IConversationRepo.
func (c *conversationRepo) DissovelGroupConversation(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "DissovelGroupConversation - "
	var query string = "Update " + business_object.GetConversationTable() + " set status = false and updated_at = ? where id = ?"
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
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetConversationTable()))
	}

	return nil
}

// GetAllConversations implements repo.IConversationRepo.
func (c *conversationRepo) GetAllConversations(ctx context.Context) (*[]businessobject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "GetAllConversations - "
	var query string = "Select * from " + business_object.GetConversationTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Conversation
	for rows.Next() {
		var x business_object.Conversation

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &x.IsDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetConversation implements repo.IConversationRepo.
func (c *conversationRepo) GetConversation(id string, ctx context.Context) (*businessobject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "GetConversation - "
	var query string = "Select * from " + business_object.GetConversationTable() + " where id = ?"

	defer c.db.Close()

	var res *business_object.Conversation
	if err := c.db.QueryRow(query, id).Scan(&res.ConversationId, &res.ConversationName, &res.HostId, &res.Members, &res.IsGroup, &res.IsDelete, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		c.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetConversationsByKeyword implements repo.IConversationRepo.
func (c *conversationRepo) GetConversationsByKeyword(id string, keyword string, ctx context.Context) (*[]businessobject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "GetConversationsFromUser - "
	var query string = "Select * from " + business_object.GetConversationTable() + " where members like '%?%' and lower(conversation_name) like lower('%?%')"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query, id, keyword)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Conversation
	for rows.Next() {
		var x business_object.Conversation

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &x.IsDelete, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		var count int = 0
		for _, name := range util.ToSliceString(x.ConversationName, "|") {
			if strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
				count += 1

				if count == 2 {
					*res = append(*res, x)
					break
				}
			}
		}
	}

	return res, nil
}

// GetConversationsFromUser implements repo.IConversationRepo.
func (c *conversationRepo) GetConversationsFromUser(id string, ctx context.Context) (*[]businessobject.Conversation, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetConversationTable()) + "GetConversationsFromUser - "
	var query string = "Select * from " + business_object.GetConversationTable() + " where members like '%?%'"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query, id)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Conversation
	for rows.Next() {
		var x business_object.Conversation

		if err := rows.Scan(&x.ConversationId, &x.ConversationName, &x.HostId, &x.Members, &x.IsGroup, &x.IsDelete, &x.CreatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}
