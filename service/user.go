package service

import (
	"context"
	"errors"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
)

type userService struct {
	logger   *log.Logger
	userRepo repo.IUserRepo
}

const (
	sepChar string = "|"
)

func GenerateUserService() (service.IUserService, error) {
	db, err := db.ConnectDB(business_object.GetUserTable())

	if err != nil {
		return nil, err
	}

	var logger *log.Logger = &log.Logger{}

	return &userService{
		logger:   logger,
		userRepo: repository.InitializeUserRepo(db, logger),
	}, nil
}

// ChangeUserStatus implements service.IUserService.
func (u *userService) ChangeUserStatus(rawStatus string, userId string, actorId string, c context.Context) (error, string) {
	panic("unimplemented")
}

// GetAllUsers implements service.IUserService.
func (u *userService) GetAllUsers(ctx context.Context) (*[]business_object.User, error) {
	tmpStorage, err := u.userRepo.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// GetUsersByRole implements service.IUserService.
func (u *userService) GetUsersByRole(role string, ctx context.Context) (*[]business_object.User, error) {
	role = util.ToNormalizedString(role)

	var tmpStorage *[]dto.UserDBResModel
	var err error

	if role == "" {
		tmpStorage, err = u.userRepo.GetAllUsers(ctx)
	} else {
		tmpStorage, err = u.userRepo.GetUsersByRole(role, ctx)
	}

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// GetUsersByStatus implements service.IUserService.
func (u *userService) GetUsersByStatus(rawStatus string, ctx context.Context) (*[]business_object.User, error) {
	rawStatus = util.ToNormalizedString(rawStatus)

	var tmpStorage *[]dto.UserDBResModel
	var err error

	if rawStatus == "" {
		tmpStorage, err = u.userRepo.GetAllUsers(ctx)
	} else {
		status, errRes := util.ToBoolean(rawStatus)

		if errRes != nil {
			return nil, errRes
		}

		tmpStorage, err = u.userRepo.GetUsersByStatus(status, ctx)
	}

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// LogOut implements service.IUserService.
func (u *userService) LogOut(id string, ctx context.Context) error {
	panic("unimplemented")
}

// Login implements service.IUserService.
func (u *userService) Login(req dto.LoginRequest, ctx context.Context) (string, string, error) {
	panic("unimplemented")
}

// CreateUser implements service.IUserService.
func (u *userService) CreateUser(req dto.CreateUserReq, ctx context.Context) error {
	panic("unimplemented")
}

// GetUser implements service.IUserService.
func (u *userService) GetUser(id string, ctx context.Context) (*business_object.User, error) {
	if id == "" {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	user, err := u.userRepo.GetUser(id, ctx)
	if err != nil {
		return nil, err
	}

	var res = toUserModel(*user)
	return &res, nil
}

func toSliceUserModel(src *[]dto.UserDBResModel) *[]business_object.User {
	var res *[]business_object.User

	for _, user := range *src {
		*res = append(*res, toUserModel(user))
	}

	return res
}

func toUserModel(src dto.UserDBResModel) business_object.User {
	var friends = util.ToSliceString(src.Friends, sepChar)
	var followers = util.ToSliceString(src.Followers, sepChar)
	var followings = util.ToSliceString(src.Followings, sepChar)
	var blockUsers = util.ToSliceString(src.BlockUsers, sepChar)

	return business_object.User{
		UserId:        src.UserId,
		RoleId:        src.RoleId,
		Username:      src.Username,
		Email:         src.Email,
		Password:      src.Password,
		DateOfBirth:   src.DateOfBirth,
		ProfileAvatar: src.ProfileAvatar,
		Bio:           src.Bio,
		IsPrivate:     &src.IsPrivate,
		IsActive:      &src.IsActive,
		Friends:       &friends,
		Followers:     &followers,
		Followings:    &followings,
		BlockUsers:    &blockUsers,
		CreatedAt:     src.CreatedAt,
		UpdatedAt:     src.UpdatedAt,
	}
}
