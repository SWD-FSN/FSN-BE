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

type userSecurityRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeUserSecurityRepo(db *sql.DB, logger *log.Logger) repo.IUserSecurityRepo {
	return &userSecurityRepo{
		db:     db,
		logger: logger,
	}
}

// EditUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) EditUserSecurity(usc business_object.UserSecurity, ctx context.Context) error {
	var query string = "Update " + business_object.GetUserSecurityTable() + " set access_token = ?, refresh_token = ?, action_token = ?, fail_access = ?, last_fail = ? where id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "EditUserSecurity - "
	var internalErr error = errors.New(noti.InternalErr)
	defer u.db.Close()

	res, err := u.db.Exec(query, usc.AccessToken, usc.RefreshToken, usc.RefreshToken, usc.ActionToken, usc.FailAccess, &usc.LastFail, usc.UserId)
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
	var query string = "Select top 1 * from " + business_object.GetUserSecurityTable() + "Where id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "GetUserSecurity - "
	defer u.db.Close()

	var res *business_object.UserSecurity
	if err := u.db.QueryRow(query, id).Scan(&res.UserId, &res.AccessToken, &res.RefreshToken, &res.ActionToken, &res.FailAccess, &res.LastFail); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}
