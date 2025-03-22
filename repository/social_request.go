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

type socialRequestRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeSocialRequestRepo(db *sql.DB, logger *log.Logger) repo.ISocialRequestRepo {
	return &socialRequestRepo{
		db:     db,
		logger: logger,
	}
}

// CreateRequest implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) CreateRequest(req business_object.SocialRequest, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "CreateAddFrRq - "
	var query string = "INSERT INTO " + business_object.GetSocialRequestTable() + "(id, author_id, account_id, request_type, created_at) VALUES ($1, $2, $3, $4, $5)"

	defer a.db.Close()

	if _, err := a.db.Exec(query, req.RequestId, req.AuthorId, req.AccountId, req.RequestType, req.CreatedAt); err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetRequest implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) GetRequest(id string, ctx context.Context) (*business_object.SocialRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "GetRequest - "
	var query string = "SELECT * FROM " + business_object.GetSocialRequestTable() + " WHERE id = $1"

	defer a.db.Close()

	var res *business_object.SocialRequest
	if err := a.db.QueryRow(query, id).Scan(&res.RequestId, &res.AuthorId, &res.AccountId, &res.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		a.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetRequestsToUser implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) GetRequestsToUser(id string, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "GetRequestsToUser - "
	var query string = "SELECT * FROM " + business_object.GetSocialRequestTable() + " WHERE request_type = $1 AND account_id = $2"
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query, requestType, id)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.SocialRequest
	for rows.Next() {
		var x business_object.SocialRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.RequestType, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetAllRequests implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) GetAllRequests(ctx context.Context) (*[]business_object.SocialRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "GetAllRequests - "
	var query string = "SELECT * FROM " + business_object.GetSocialRequestTable()
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.SocialRequest
	for rows.Next() {
		var x business_object.SocialRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.RequestType, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetUserRequests implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) GetUserRequests(id string, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "GetUserRequests - "
	var query string = "SELECT * FROM " + business_object.GetSocialRequestTable() + " WHERE request_type = $1 AND author_id = $2"
	var internalErr error = errors.New(noti.InternalErr)

	defer a.db.Close()

	rows, err := a.db.Query(query, requestType, id)
	if err != nil {
		a.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res *[]business_object.SocialRequest
	for rows.Next() {
		var x business_object.SocialRequest

		if err := rows.Scan(&x.RequestId, &x.AuthorId, &x.AccountId, &x.RequestType, &x.CreatedAt); err != nil {
			a.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// RemoveRequest implements repo.ISocialRequestRepo.
func (a *socialRequestRepo) RemoveRequest(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetSocialRequestTable()) + "RemoveRequest - "
	var query string = "DELETE FROM " + business_object.GetSocialRequestTable() + " WHERE id = $1"
	var internalErrMsg error = errors.New(noti.InternalErr)

	defer a.db.Close()

	res, err := a.db.Exec(query, id)
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
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetSocialRequestTable()))
	}

	return nil
}
