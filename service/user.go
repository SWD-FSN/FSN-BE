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
