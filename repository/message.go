package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
)

type messageRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeMessageRepo(db *sql.DB, logger *log.Logger) repo.IMessageRepo {
	return &messageRepo{
		db:     db,
		logger: logger,
	}
}

// GetMessagesFromConversationByKeyword implements repo.IMessageRepo.
func (m *messageRepo) GetMessagesFromConversationByKeyword(id string, keyword string, ctx context.Context) (*[]business_object.Message, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "GetMessagesFROMConversationByKeyword - "
	var query string = "SELECT * FROM " + business_object.GetMessageTable() + " WHERE lower(content) LIKE lower('%$1%'), conversation_id = $2\n ORDER BY created_at DESC"
	var internalErr error = errors.New(noti.InternalErr)

	//defer m.db.Close()

	rows, err := m.db.Query(query, keyword, id)
	if err != nil {
		m.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Message
	for rows.Next() {
		var x business_object.Message

		if err := rows.Scan(&x.MessageId, &x.AuthorId, &x.CoversationId, &x.Content, &x.CreatedAt); err != nil {
			m.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// CreateMessage implements repo.IMessageRepo.
func (m *messageRepo) CreateMessage(message business_object.Message, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "CreateMessage - "
	var query string = "INSERT INTO " + business_object.GetMessageTable() + "(id, author_id, conversation_id, content, created_at) values ($1, $2, $3, $4, $5)"

	//defer m.db.Close()

	if _, err := m.db.Exec(query, message.MessageId, message.AuthorId, message.CoversationId, message.Content, message.CreatedAt); err != nil {
		m.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAllMessages implements repo.IMessageRepo.
func (m *messageRepo) GetAllMessages(ctx context.Context) (*[]business_object.Message, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "GetAllMessages - "
	var query string = "SELECT * FROM " + business_object.GetMessageTable()
	var internalErr error = errors.New(noti.InternalErr)

	//defer m.db.Close()

	rows, err := m.db.Query(query)
	if err != nil {
		m.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Message
	for rows.Next() {
		var x business_object.Message

		if err := rows.Scan(&x.MessageId, &x.AuthorId, &x.CoversationId, &x.Content, &x.CreatedAt); err != nil {
			m.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetMessage implements repo.IMessageRepo.
func (m *messageRepo) GetMessage(id string, ctx context.Context) (*business_object.Message, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "GetMessage - "
	var query string = "SELECT * FROM " + business_object.GetMessageTable() + " WHERE id = $1"

	//defer m.db.Close()

	var res business_object.Message
	if err := m.db.QueryRow(query, id).Scan(&res.MessageId, &res.AuthorId, &res.CoversationId, &res.Content, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		m.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return &res, nil
}

// GetMessagesFROMConversation implements repo.IMessageRepo.
func (m *messageRepo) GetMessagesFromConversation(id string, ctx context.Context) (*[]business_object.Message, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "GetAllMessages - "
	var query string = "SELECT * FROM " + business_object.GetMessageTable() + " WHERE conversation_id = $1\n ORDER BY created_at DESC"
	var internalErr error = errors.New(noti.InternalErr)

	//defer m.db.Close()

	rows, err := m.db.Query(query, id)
	if err != nil {
		m.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Message
	for rows.Next() {
		var x business_object.Message

		if err := rows.Scan(&x.MessageId, &x.AuthorId, &x.CoversationId, &x.Content, &x.CreatedAt); err != nil {
			m.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetMessagesFromUser implements repo.IMessageRepo.
func (m *messageRepo) GetMessagesFromUser(id string, ctx context.Context) (*[]business_object.Message, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetMessageTable()) + "GetAllMessages - "
	var query string = "SELECT * FROM " + business_object.GetMessageTable() + " WHERE author_id = $1"
	var internalErr error = errors.New(noti.InternalErr)

	//defer m.db.Close()

	rows, err := m.db.Query(query, id)
	if err != nil {
		m.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Message
	for rows.Next() {
		var x business_object.Message

		if err := rows.Scan(&x.MessageId, &x.AuthorId, &x.CoversationId, &x.Content, &x.CreatedAt); err != nil {
			m.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}
