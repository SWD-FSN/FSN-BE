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

type likeRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeLikeRepo(db *sql.DB, logger *log.Logger) repo.ILikeRepo {
	return &likeRepo{
		db:     db,
		logger: logger,
	}
}

// CancelLike implements repo.ILikeRepo.
func (l *likeRepo) CancelLike(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "CancelLike - "
	var query string = "Delete from " + business_object.GetLikeTable() + " where id = ?"

	defer l.db.Close()

	res, err := l.db.Exec(query, id)
	var internalErrMsg error = errors.New(noti.InternalErr)

	if err != nil {
		l.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		l.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetLikeTable()))
	}

	return nil
}

// CreateLike implements repo.ILikeRepo.
func (l *likeRepo) CreateLike(like business_object.Like, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "CreateLike - "
	var query string = "Insert into " + business_object.GetLikeTable() + "(id, author_id, object_id, object_type, created_at) values (?, ?, ?, ?, ?)"

	defer l.db.Close()

	if _, err := l.db.Exec(query, like.LikeId, like.AuthorId, like.ObjectId, like.ObjectType, like.CreatedAt); err != nil {
		l.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAllLikes implements repo.ILikeRepo.
func (l *likeRepo) GetAllLikes(ctx context.Context) (*[]business_object.Like, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "GetAllLikes - "
	var query string = "Select * from " + business_object.GetLikeTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer l.db.Close()

	rows, err := l.db.Query(query)
	if err != nil {
		l.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Like
	for rows.Next() {
		var x business_object.Like

		if err := rows.Scan(&x.LikeId, &x.AuthorId, &x.ObjectId, &x.ObjectId, &x.ObjectType, &x.CreatedAt); err != nil {
			l.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetLike implements repo.ILikeRepo.
func (l *likeRepo) GetLike(id string, ctx context.Context) (*business_object.Like, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "GetLike - "
	var query string = "Select * from " + business_object.GetLikeTable() + " where id = ?"

	defer l.db.Close()

	var res *business_object.Like
	if err := l.db.QueryRow(query, id).Scan(&res.LikeId, &res.AuthorId, &res.ObjectId, &res.ObjectType, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		l.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetLikesFromObject implements repo.ILikeRepo.
func (l *likeRepo) GetLikesFromObject(id string, kind string, ctx context.Context) (*[]business_object.Like, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "GetLikesFromObject - "
	var query string = "Select * from " + business_object.GetLikeTable() + " where object_id = ? and object_type = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer l.db.Close()

	rows, err := l.db.Query(query, id, kind)
	if err != nil {
		l.logger.Println(errLogMsg, err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Like
	for rows.Next() {
		var x business_object.Like

		if err := rows.Scan(&x.LikeId, &x.AuthorId, &x.ObjectId, &x.ObjectId, &x.ObjectType, &x.CreatedAt); err != nil {
			l.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}
