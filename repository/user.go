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
	"social_network/util"
	"sync"
	"time"
)

type userRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeUserRepo(db *sql.DB, logger *log.Logger) repo.IUserRepo {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

const (
	friends_involed    string = "FRIENDS_INVOLED"
	blocks_involed     string = "BLOCKEDS_INVOLED"
	followers_involed  string = "FOLLOWERS_INVOLED"
	followings_involed string = "FOLLOWINGS_INVOLED"
)

// ChangeUserStatus implements repo.IUserRepo.
func (u *userRepo) ChangeUserStatus(id string, status bool, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "ChangeUserStatus - "

	var lastFailValueQuery string = "NULL"
	if status {
		lastFailValueQuery = fmt.Sprint(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))
	}

	var userQuery string = "UPDATE " + business_object.GetUserTable() + " SET is_active = $1 AND updated_at = $2 WHERE user_id = $3"
	var securityQuery string = "UPDATE " + business_object.GetUserSecurityTable() + " SET access_token = NULL, access_expiration = NULL," +
		" refresh_token = NULL, refresh_expiration = NULL, action_token = NULL, action_expiration = NULL, fail_access = 0, last_fail = " +
		lastFailValueQuery + " WHERE id = $1"
	defer u.db.Close()

	var errChan chan error = make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := u.db.Exec(userQuery, fmt.Sprint(status), time.Now().GoString(), id); err != nil {
			u.logger.Println(errLogMsg + err.Error())
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := u.db.Exec(securityQuery, id); err != nil {
			u.logger.Println(errLogMsg + err.Error())
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return errors.New(noti.InternalErr)
		}
	}

	return nil
}

// GetAllUsers implements repo.IUserRepo.
func (u *userRepo) GetAllUsers(ctx context.Context) (*[]dto.UserDBResModel, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetAllUsers - "
	var query string = "SELECT * FROM " + business_object.GetUserTable()
	defer u.db.Close()

	rows, err := u.db.Query(query)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res []dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		var isActivated sql.NullBool
		var isHaveToResetPw sql.NullBool

		if err := rows.Scan(
			&x.UserId, &x.RoleId, &x.FullName, &x.Username, &x.Email, &x.Password,
			&x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Friends, &x.Followers,
			&x.Followings, &x.BlockUsers, &x.Conversations, &x.IsPrivate,
			&x.IsActive, &isActivated, &isHaveToResetPw, &x.CreatedAt, &x.UpdatedAt); err != nil {

			u.logger.Println(errLogMsg + err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		x.IsActivated = isActivated.Valid && isActivated.Bool
		if isHaveToResetPw.Valid {
			boolVal := isHaveToResetPw.Bool
			x.IsHaveToResetPw = &boolVal
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetUsersByStatus implements repo.IUserRepo.
func (u *userRepo) GetUsersByStatus(status bool, ctx context.Context) (*[]dto.UserDBResModel, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUsersByStatus - "
	var query string = "SELECT * FROM " + business_object.GetUserTable() + " WHERE is_active = $1"
	defer u.db.Close()

	rows, err := u.db.Query(query, status)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res []dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		var isActivated sql.NullBool
		var isHaveToResetPw sql.NullBool

		if err := rows.Scan(
			&x.UserId, &x.RoleId, &x.FullName, &x.Username, &x.Email, &x.Password,
			&x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Friends, &x.Followers,
			&x.Followings, &x.BlockUsers, &x.Conversations, &x.IsPrivate,
			&x.IsActive, &isActivated, &isHaveToResetPw, &x.CreatedAt, &x.UpdatedAt); err != nil {

			u.logger.Println(errLogMsg + err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		x.IsActivated = isActivated.Valid && isActivated.Bool
		if isHaveToResetPw.Valid {
			boolVal := isHaveToResetPw.Bool
			x.IsHaveToResetPw = &boolVal
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetUserByEmail implements repo.IUserRepo.
func (u *userRepo) GetUserByEmail(email string, ctx context.Context) (*dto.UserDBResModel, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUserByEmail - "
	var query string = "SELECT * from " + business_object.GetUserTable() + " WHERE LOWER(email) = LOWER($1)"
	//defer u.db.Close()

	var res dto.UserDBResModel
	var isActivated sql.NullBool
	var isHaveToResetPw sql.NullBool

	if err := u.db.QueryRow(query, email).Scan(
		&res.UserId, &res.RoleId, &res.FullName, &res.Username, &res.Email, &res.Password,
		&res.DateOfBirth, &res.ProfileAvatar, &res.Bio, &res.Friends, &res.Followers,
		&res.Followings, &res.BlockUsers, &res.Conversations, &res.IsPrivate,
		&res.IsActive, &isActivated, &isHaveToResetPw, &res.CreatedAt, &res.UpdatedAt); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	// Transfer nullable values to the result struct
	res.IsActivated = isActivated.Valid && isActivated.Bool
	if isHaveToResetPw.Valid {
		boolVal := isHaveToResetPw.Bool
		res.IsHaveToResetPw = &boolVal
	}

	return &res, nil
}

// GetUsersByRole implements repo.IUserRepo.
func (u *userRepo) GetUsersByRole(id string, ctx context.Context) (*[]dto.UserDBResModel, error) {
	var query string = "SELECT * from " + business_object.GetUserTable() + " WHERE role_id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUsersByRole - "
	defer u.db.Close()

	rows, err := u.db.Query(query, id)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res []dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		var isActivated sql.NullBool
		var isHaveToResetPw sql.NullBool

		if err := rows.Scan(
			&x.UserId, &x.RoleId, &x.FullName, &x.Username, &x.Email, &x.Password,
			&x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Friends, &x.Followers,
			&x.Followings, &x.BlockUsers, &x.Conversations, &x.IsPrivate,
			&x.IsActive, &isActivated, &isHaveToResetPw, &x.CreatedAt, &x.UpdatedAt); err != nil {

			u.logger.Println(errLogMsg + err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		x.IsActivated = isActivated.Valid && isActivated.Bool
		if isHaveToResetPw.Valid {
			boolVal := isHaveToResetPw.Bool
			x.IsHaveToResetPw = &boolVal
		}

		res = append(res, x)
	}

	return &res, nil
}

// UpdateUser implements repo.IUserRepo.
func (u *userRepo) UpdateUser(user dto.UserDBResModel, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "UpdateUser - "
	var query string = "UPDATE " + business_object.GetUserTable() + " SET email = $1, password = $2, role_id = $3, " +
		"full_name = $4, username = $5, date_of_birth = $6, profile_avatar = $7, bio = $8, followers = $9, " +
		"followings = $10, block_users = $11, conversations = $12, is_private = $13, is_active = $14, is_activated = $15, " +
		"updated_at = $16 WHERE id = $17"
	defer u.db.Close()

	res, err := u.db.Exec(query, user.Email, user.Password, user.RoleId, user.FullName, user.Username, user.DateOfBirth,
		user.ProfileAvatar, user.Bio, user.Followers, user.Followings, user.BlockUsers,
		user.Conversations, user.IsPrivate, user.IsActive, user.IsActivated,
		time.Now(), user.UserId)

	var internalErrMsg error = errors.New(noti.InternalErr)

	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetUserTable()))
	}

	return nil
}

// CreateUser implements repo.IUserRepo.
func (u *userRepo) CreateUser(user dto.UserDBResModel, ctx context.Context) error {
	var query string = "INSERT INTO " + business_object.GetUserTable() +
		"(id, role_id, full_name, username, email, password, date_of_birth, " +
		"profile_avatar, bio, friends, followers, followings, block_users, conversations, " +
		"is_private, is_active, created_at, updated_at) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "CreateUser - "

	//defer u.db.Close()

	if _, err := u.db.Exec(query, user.UserId, user.RoleId, user.FullName, user.Username, user.Email, user.Password, user.DateOfBirth, user.ProfileAvatar, user.Bio, user.Friends, user.Followers, user.Followings, user.BlockUsers, user.Conversations, user.IsPrivate, user.IsActive, user.CreatedAt, user.UpdatedAt); err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetUser implements repo.IUserRepo.
func (u *userRepo) GetUser(id string, ctx context.Context) (*dto.UserDBResModel, error) {
	var query string = "SELECT * FROM " + business_object.GetUserTable() + " WHERE id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUser - "
	//defer u.db.Close()

	var res dto.UserDBResModel
	var isActivated sql.NullBool
	var isHaveToResetPw sql.NullBool

	var err error = u.db.QueryRow(query, id).Scan(
		&res.UserId, &res.RoleId, &res.FullName, &res.Username, &res.Email, &res.Password,
		&res.DateOfBirth, &res.ProfileAvatar, &res.Bio, &res.Friends, &res.Followers,
		&res.Followings, &res.BlockUsers, &res.Conversations, &res.IsPrivate,
		&res.IsActive, &isActivated, &isHaveToResetPw, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	res.IsActivated = isActivated.Valid && isActivated.Bool
	if isHaveToResetPw.Valid {
		boolVal := isHaveToResetPw.Bool
		res.IsHaveToResetPw = &boolVal
	}

	return &res, nil
}

// GetInvoledAccountsAmountFromUser implements repo.IUserRepo.
func (u *userRepo) GetInvolvedAccountsAmountFromUser(req dto.GetInvoledAccouuntsRequest, ctx context.Context) ([]string, error) {
	var field string = getFieldFromInvolvedRequest(req.InvolvedType)
	if field == "" {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	var query string = "SELECT " + field + " FROM " + business_object.GetUserTable() + " WHERE id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetInvolvedAccountsAmountFromUser - "
	defer u.db.Close()

	var combinedString string
	if err := u.db.QueryRow(query, req.UserId).Scan(&combinedString); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return util.ToSliceString(combinedString, "|"), nil
}

// GetInvolvedAccountsFromTag implements repo.IUserRepo.
func (u *userRepo) GetInvolvedAccountsFromTag(id string, ctx context.Context) ([]string, error) {
	var query string = "SELECT friends, followers, followings FROM " + business_object.GetUserTable() + " WHERE id = $1"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetInvolvedAccountsFromTag - "

	defer u.db.Close()

	var friends, followers, followings string

	if err := u.db.QueryRow(query, id).Scan(&friends, &followers, &followings); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return util.ToSliceString(friends+"|"+followers+"|"+followings, "|"), nil
}

// GetUsersByKeyword implements repo.IUserRepo.
func (u *userRepo) GetUsersByKeyword(keyword string, ctx context.Context) (*[]dto.UserDBResModel, error) {
	var query string = "SELECT * FROM " + business_object.GetUserTable() +
		" WHERE LOWER(username) LIKE LOWER('%$1%') OR LOWER(full_name) LIKE ('%$2%') OR LOWER(email) LIKE ('%$3%')"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUsersByKeyword - "
	var internalErr error = errors.New(noti.InternalErr)

	defer u.db.Close()

	rows, err := u.db.Query(query, keyword, keyword, keyword)
	if err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		var isActivated sql.NullBool
		var isHaveToResetPw sql.NullBool

		if err := rows.Scan(
			&x.UserId, &x.RoleId, &x.FullName, &x.Username, &x.Email, &x.Password,
			&x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Friends, &x.Followers,
			&x.Followings, &x.BlockUsers, &x.Conversations, &x.IsPrivate,
			&x.IsActive, &isActivated, &isHaveToResetPw, &x.CreatedAt, &x.UpdatedAt); err != nil {

			u.logger.Println(errLogMsg + err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		x.IsActivated = isActivated.Valid && isActivated.Bool
		if isHaveToResetPw.Valid {
			boolVal := isHaveToResetPw.Bool
			x.IsHaveToResetPw = &boolVal
		}

		res = append(res, x)
	}

	return &res, nil
}

func getFieldFromInvolvedRequest(req string) string {
	var res string

	switch req {
	case friends_involed:
		res = "friends"
	case followers_involed:
		res = "followers"
	case followings_involed:
		res = "followings"
	case blocks_involed:
		res = "block_users"
	}

	return res
}
