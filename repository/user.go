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

// CreateUser implements repo.IUserRepo.
func (u *userRepo) CreateUser(user dto.UserSaveModel, ctx context.Context) error {
	var query string = "Insert into " + business_object.GetUserTable() + "(username, phone_number, date_of_birth, profile_avatar, bio, followers, followings, block_users, conversations, is_private, is_active, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "CreateUser - "

	defer u.db.Close()

	if _, err := u.db.Exec(query, user.UserName, user.PhoneNumber, user.DateOfBirth, user.ProfileAvatar, user.Bio, user.Followers, user.Followings, user.BlockUsers, user.Conversations, user.IsPrivate, user.IsActive, user.CreatedAt, user.UpdatedAt); err != nil {
		u.logger.Println(errLogMsg, err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetUser implements repo.IUserRepo.
func (u *userRepo) GetUser(id string, ctx context.Context) (*dto.UserDBResModel, error) {
	var query string = "Select username, phone_number, date_of_birth, profile_avatar, bio, followers, followings, block_users, conversations, is_private, is_active, created_at, updated_at from " + business_object.GetUserTable() + "Where id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserTable()) + "GetUser - "

	defer u.db.Close()

	var res *dto.UserDBResModel

	var err error = u.db.QueryRow(query, id).Scan(&res.UserName, &res.PhoneNumber, &res.DateOfBirth, &res.ProfileAvatar, &res.Bio, &res.Followers, &res.Followings, &res.BlockUsers, &res.Conversations, &res.IsPrivate, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		u.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	res.UserId = id
	return res, nil
}
