package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/dto"
	"social_network/interfaces/repo"
)

type userSecurityRepo struct {
	db     *sql.DB
	logger *log.Logger
}

// CreateUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) CreateUserSecurity(usc business_object.UserSecurity, ctx context.Context) error {
	var query string = "INSERT INTO " + business_object.GetUserSecurityTable() + " (user_id, access_token, refresh_token, action_token, fail_access, last_fail) VALUES ($1, $2, $3, $4, $5, $6)"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "CreateUserSecurity - "

	if _, err := u.db.Exec(query, usc.UserId, usc.AccessToken, usc.RefreshToken, usc.ActionToken, usc.FailAccess, usc.LastFail); err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

func InitializeUserSecurityRepo(db *sql.DB, logger *log.Logger) repo.IUserSecurityRepo {
	return &userSecurityRepo{
		db:     db,
		logger: logger,
	}
}

// Login implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) Login(req dto.LoginSecurityRequest, ctx context.Context) error {
	var query string = "UPDATE " + business_object.GetUserSecurityTable() + " SET access_token = $1, refresh_token = $2 WHERE user_id = $3"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "Login - "
	var internalErr error = errors.New(noti.InternalErr)

	res, err := u.db.Exec(query, req.AccessToken, req.RefreshToken, req.UserId)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetUserSecurityTable()))
	}

	return nil
}

// Logout implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) Logout(id string, ctx context.Context) error {
	var query string = "UPDATE " + business_object.GetUserSecurityTable() + " SET access_token = NULL, refresh_token = NULL WHERE user_id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "LogOut - "
	var internalErr error = errors.New(noti.InternalErr)
	//defer u.db.Close()

	res, err := u.db.Exec(query, id)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetUserSecurityTable()))
	}

	return nil
}

// EditUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) EditUserSecurity(usc business_object.UserSecurity, ctx context.Context) error {
	var query string = "UPDATE " + business_object.GetUserSecurityTable() + " SET access_token = $1, refresh_token = $2, action_token = $3, fail_access = $4, last_fail = $5 WHERE user_id = $6"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "EditUserSecurity - "
	var internalErr error = errors.New(noti.InternalErr)

	res, err := u.db.Exec(query, usc.AccessToken, usc.RefreshToken, usc.ActionToken, usc.FailAccess, &usc.LastFail, usc.UserId)
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetUserSecurityTable()))
	}

	return nil
}

// GetUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) GetUserSecurity(id string, ctx context.Context) (*business_object.UserSecurity, error) {
	var query string = "SELECT * FROM " + business_object.GetUserSecurityTable() + " WHERE user_id = $1 LIMIT 1"

	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "GetUserSecurity - "

	var usc business_object.UserSecurity
	if err := u.db.QueryRow(query, id).Scan(
		&usc.UserId, &usc.AccessToken, &usc.RefreshToken,
		&usc.ActionToken, &usc.FailAccess, &usc.LastFail); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}
	return &usc, nil
}
