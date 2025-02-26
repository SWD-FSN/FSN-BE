package service

import (
	"context"
	"errors"
	"log"
	businessobject "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
)

type userSecurityService struct {
	userSecurityRepo repo.IUserSecurityRepo
	logger           *log.Logger
}

func GenerateUserSecurityService() (service.IUserSecurityService, error) {
	var logger = util.GetLogConfig()

	db, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return &userSecurityService{
		userSecurityRepo: repository.InitializeUserSecurityRepo(db, logger),
		logger:           logger,
	}, nil
}

// EditUserSecurity implements service.IUserSecurityService.
func (u *userSecurityService) EditUserSecurity(usc businessobject.UserSecurity, ctx context.Context) error {
	if usc.UserId == "" {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	return u.userSecurityRepo.EditUserSecurity(usc, ctx)
}

// GetUserSecurity implements service.IUserSecurityService.
func (u *userSecurityService) GetUserSecurity(id string, ctx context.Context) (*businessobject.UserSecurity, error) {
	if id == "" {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	return u.userSecurityRepo.GetUserSecurity(id, ctx)
}
