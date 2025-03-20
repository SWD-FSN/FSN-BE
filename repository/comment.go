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

type commentRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeCommentRepo(db *sql.DB, logger *log.Logger) repo.ICommentRepo {
	return &commentRepo{
		db:     db,
		logger: logger,
	}
}

// CreateComment implements repo.ICommentRepo.
func (c *commentRepo) CreateComment(cmt business_object.Comment, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetCommentTable()) + "CreateComment - "
	var query string = "INSERT INTO " + business_object.GetCommentTable() + "(id, author_id, post_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	defer c.db.Close()

	if _, err := c.db.Exec(query, cmt.CommentId, cmt.AuthorId, cmt.PostId, cmt.Content, cmt.CreatedAt, cmt.UpdatedAt); err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// EditComment implements repo.ICommentRepo.
func (c *commentRepo) EditComment(cmt business_object.Comment, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetCommentTable()) + "EditComment - "
	var query string = "UPDATE " + business_object.GetPostTable() + " SET content = ? AND updated_at = ? WHERE id = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	res, err := c.db.Exec(query, cmt.Content, cmt.UpdatedAt, cmt.CommentId)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetCommentTable()))
	}

	return nil
}

// GetCommentsFromPost implements repo.ICommentRepo.
func (c *commentRepo) GetCommentsFromPost(id string, ctx context.Context) (*[]business_object.Comment, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetCommentTable()) + "GetCommentsFromPost - "
	var query string = "SELECT * FROM " + business_object.GetCommentTable() + " WHERE post_id = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer c.db.Close()

	rows, err := c.db.Query(query, id)
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Comment
	for rows.Next() {
		var x business_object.Comment

		if err := rows.Scan(&x.CommentId, &x.AuthorId, &x.PostId, &x.Content, &x.CreatedAt, &x.UpdatedAt); err != nil {
			c.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetComment implements repo.ICommentRepo.
func (c *commentRepo) GetComment(id string, ctx context.Context) (*business_object.Comment, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetCommentTable()) + "GetComment - "
	var query string = "SELECT * FROM  " + business_object.GetCommentTable() + " WHERE id = ?"

	defer c.db.Close()

	var res *business_object.Comment
	if err := c.db.QueryRow(query, id).Scan(&res.CommentId, &res.PostId, &res.AuthorId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		c.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// RemoveComment implements repo.ICommentRepo.
func (c *commentRepo) RemoveComment(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetCommentTable()) + "RemoveComment - "
	var query string = "DELETE FROM " + business_object.GetCommentTable() + " WHERE id = ?"

	defer c.db.Close()

	res, err := c.db.Exec(query, id)
	var internalErrMsg error = errors.New(noti.InternalErr)

	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetLikeTable()))
	}

	return nil
}
