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

// ChangeUserStatus implements repo.IUserRepo.
func (u *userRepo) ChangeUserStatus(id string, status bool, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "ChangeUserStatus - "

	var lastFailValueQuery string = "NULL"
	if status {
		lastFailValueQuery = fmt.Sprint(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))
	}

	var userQuery string = "Update " + business_object.GetUserTable() + " set is_active = ?, updated_at = ? where id = ?"
	var securityQuery string = "Update " + business_object.GetUserSecurityTable() + " set access_token = NULL, access_expiration = NULL, refresh_token = NULL, refresh_expiration = NULL, action_token = NULL, action_expiration = NULL, fail_access = 0, last_fail = " + lastFailValueQuery + " where id = ?"
	defer u.db.Close()

	var errChan chan error = make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := u.db.Exec(userQuery, fmt.Sprint(status), time.Now().UTC().GoString(), id); err != nil {
			u.logger.Println(errLogMsg, err.Error())
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := u.db.Exec(securityQuery, id); err != nil {
			u.logger.Println(errLogMsg, err.Error())
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
	var query string = "Select * from " + business_object.GetUserTable()
	defer u.db.Close()

	rows, err := u.db.Query(query)
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res *[]dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		if err := rows.Scan(&x.UserId, &x.RoleId, &x.Username, &x.Email, &x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Followers, &x.Followings, &x.BlockUsers, &x.Conversations, &x.IsActive, &x.IsActive, &x.CreatedAt, &x.UpdatedAt); err != nil {
			u.logger.Println(errLogMsg, err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetUsersByStatus implements repo.IUserRepo.
func (u *userRepo) GetUsersByStatus(status bool, ctx context.Context) (*[]dto.UserDBResModel, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUsersByStatus - "
	var query string = "Select * from " + business_object.GetUserTable() + "where is_active = ?"
	defer u.db.Close()

	rows, err := u.db.Query(query, status)
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res *[]dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		if err := rows.Scan(&x.UserId, &x.RoleId, &x.Username, &x.Email, &x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Followers, &x.Followings, &x.BlockUsers, &x.Conversations, &x.IsActive, &x.IsActive, &x.CreatedAt, &x.UpdatedAt); err != nil {
			u.logger.Println(errLogMsg, err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetUserByEmail implements repo.IUserRepo.
func (u *userRepo) GetUserByEmail(email string, ctx context.Context) (*dto.UserDBResModel, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUserByEmail - "
	var query string = "Select top 1 * from " + business_object.GetUserTable() + "where lower(email) = lower($1)"
	defer u.db.Close()

	var res *dto.UserDBResModel
	if err := u.db.QueryRow(query, email).Scan(&res.UserId, &res.RoleId, &res.Username, &res.Email, &res.DateOfBirth, &res.ProfileAvatar, &res.Bio, &res.Followers, &res.Followings, &res.BlockUsers, &res.Conversations, &res.IsActive, &res.IsActive, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetUsersByRole implements repo.IUserRepo.
func (u *userRepo) GetUsersByRole(id string, ctx context.Context) (*[]dto.UserDBResModel, error) {
	var query string = "Select * from " + business_object.GetUserTable() + "Where role_id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUsersByRole - "
	defer u.db.Close()

	rows, err := u.db.Query(query, id)
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	var res *[]dto.UserDBResModel
	for rows.Next() {
		var x dto.UserDBResModel
		if err := rows.Scan(&x.UserId, &x.RoleId, &x.Username, &x.Email, &x.Password, &x.DateOfBirth, &x.ProfileAvatar, &x.Bio, &x.Friends, &x.Followers, &x.Followings, &x.BlockUsers, &x.Conversations, &x.IsActive, &x.IsActive, &x.CreatedAt, &x.UpdatedAt); err != nil {
			u.logger.Println(errLogMsg, err.Error())
			return nil, errors.New(noti.InternalErr)
		}

		*res = append(*res, x)
	}

	return res, nil
}

// UpdateUser implements repo.IUserRepo.
func (u *userRepo) UpdateUser(user dto.UserDBResModel) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "UpdateUser - "
	var query string = "Update " + business_object.GetUserTable() + " set email = ?, password = ?, role_id = ?, username = ?, date_of_birth = ?, profile_avatar = ?, bio = ?, followers = ?, followings = ?, block_users = ?, conversations = ?, is_private = ?, is_active = ?, updated_at = ? where id = ?"
	defer u.db.Close()

	res, err := u.db.Exec(query, user.Email, user.Password, user.RoleId, user.Username, user.DateOfBirth, user.ProfileAvatar, user.Bio, user.Followers, user.Followings, user.BlockUsers, user.Conversations, user.IsPrivate, user.IsActive, time.Now().UTC().GoString(), user.UserId)
	var internalErrMsg error = errors.New(noti.InternalErr)

	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return internalErrMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return internalErrMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetUserTable()))
	}

	return nil
}

// CreateUser implements repo.IUserRepo.
func (u *userRepo) CreateUser(user dto.UserDBResModel, ctx context.Context) error {
	var query string = "Insert into " + business_object.GetUserTable() + "(user_id, role_id, username, email, password, date_of_birth, profile_avatar, bio, friends, followers, followings, block_users, conversations, is_private, is_active, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "CreateUser - "

	defer u.db.Close()

	if _, err := u.db.Exec(query, user.UserId, user.RoleId, user.Username, user.Email, user.Password, user.DateOfBirth, user.ProfileAvatar, user.Bio, user.Friends, user.Followers, user.Followings, user.BlockUsers, user.Conversations, user.IsPrivate, user.IsActive, user.CreatedAt, user.UpdatedAt); err != nil {
		u.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetUser implements repo.IUserRepo.
func (u *userRepo) GetUser(id string, ctx context.Context) (*dto.UserDBResModel, error) {
	var query string = "Select top 1 * from " + business_object.GetUserTable() + "Where id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUser - "

	defer u.db.Close()

	var res *dto.UserDBResModel

	var err error = u.db.QueryRow(query, id).Scan(&res.Username, &res.Email, &res.DateOfBirth, &res.ProfileAvatar, &res.Bio, &res.Followers, &res.Followings, &res.BlockUsers, &res.Conversations, &res.IsPrivate, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}
