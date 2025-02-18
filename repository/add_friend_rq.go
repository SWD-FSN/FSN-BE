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

type addFriendRqRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeAddFriendRqRepo(db *sql.DB, logger *log.Logger) repo.IAddFriendRqRepo {
	return &addFriendRqRepo{
		db:     db,
		logger: logger,
	}
}

// CreateAddFrRq implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) CreateAddFrRq(req business_object.AddFrRequest, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "CreateAddFrRq - "
	var query string = "Insert into " + business_object.GetAddFriendRequestTable() + "(id, author_id, account_id, created_at) values (?, ?, ?, ?)"

	defer a.db.Close()

	if _, err := a.db.Exec(query, req.RequestId, req.AuthorId, req.AccountId, req.CreatedAt); err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAddFrRq implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) GetAddFrRq(id string, ctx context.Context) (*business_object.AddFrRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "GetAddFrRq - "
	var query string = "Select * from " + business_object.GetAddFriendRequestTable() + " where id = ?"

	defer a.db.Close()

	var res *business_object.AddFrRequest
	if err := a.db.QueryRow(query, id).Scan(&res.RequestId, &res.AuthorId, &res.AccountId, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		a.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetAddFrRqsToUser implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) GetAddFrRqsToUser(id string, ctx context.Context) (*[]business_object.AddFrRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "GetAddFrRqsToUser - "
	var query string = "Select * from " + business_object.GetAddFriendRequestTable() + " where account_id = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query, id)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.AddFrRequest
	for rows.Next() {
		var x business_object.AddFrRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetAllAddFrRqs implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) GetAllAddFrRqs(ctx context.Context) (*[]business_object.AddFrRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "GetAllAddFrRqs - "
	var query string = "Select * from " + business_object.GetAddFriendRequestTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.AddFrRequest
	for rows.Next() {
		var x business_object.AddFrRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetUserAddFrRqs implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) GetUserAddFrRqs(id string, ctx context.Context) (*[]business_object.AddFrRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "GetUserAddFrRqs - "
	var query string = "Select * from " + business_object.GetAddFriendRequestTable() + " where author_id = ?"
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query, id)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.AddFrRequest
	for rows.Next() {
		var x business_object.AddFrRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// RemoveAddFrRq implements repo.IAddFriendRqRepo.
func (a *addFriendRqRepo) RemoveAddFrRq(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetAddFriendRequestTable()) + "RemoveAddFrRq - "
	var query string = "Delete from " + business_object.GetAddFriendRequestTable() + " where id = ?"

	defer a.db.Close()

	res, err := a.db.Exec(query, id)
	var internalErrMsg error = errors.New(noti.InternalErr)

	if err != nil {
		a.logger.Println(errLogMsg, err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		a.logger.Println(errLogMsg, err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetLikeTable()))
	}

	return nil
}
