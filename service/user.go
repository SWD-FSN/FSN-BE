package service

import (
	"context"
	"errors"
	"log"
	businessobject "social_network/business_object"
	"social_network/constant/noti"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
)

type userService struct {
	logger   *log.Logger
	userRepo repo.IUserRepo
}

// ChangeUserStatus implements service.IUserService.
func (u *userService) ChangeUserStatus(rawStatus string, userId string, actorId string, c context.Context) (error, string) {
	panic("unimplemented")
}

// GetAllUsers implements service.IUserService.
func (u *userService) GetAllUsers(ctx context.Context) (*[]businessobject.User, error) {
	panic("unimplemented")
}

// GetUsersByRole implements service.IUserService.
func (u *userService) GetUsersByRole(role string, ctx context.Context) (*[]businessobject.User, error) {
	panic("unimplemented")
}

// GetUsersByStatus implements service.IUserService.
func (u *userService) GetUsersByStatus(rawStatus string, ctx context.Context) (*[]businessobject.User, error) {
	panic("unimplemented")
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
func (u *userService) GetUser(id string, ctx context.Context) (*businessobject.User, error) {
	panic("unimplemented")
}

func GenerateService() (service.IUserService, error) {
	db, err := db.ConnectDB()

	if err != nil {
		return nil, errors.New(noti.InternalErr)
	}

	var logger *log.Logger = &log.Logger{}

	return &userService{
		logger:   logger,
		userRepo: repository.InitializeUserRepo(db, logger),
	}, nil
}
