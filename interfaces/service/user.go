package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) (*[]business_object.User, error)
	GetUsersByRole(role string, ctx context.Context) (*[]business_object.User, error)
	GetUsersByStatus(rawStatus string, ctx context.Context) (*[]business_object.User, error)
	GetUser(id string, ctx context.Context) (*business_object.User, error)
	CreateUser(req dto.CreateUserReq, ctx context.Context) error
	//UpdateUser(user request.PublicUserInfo, actorId string, c context.Context) (string, error)
	ChangeUserStatus(rawStatus, userId, actorId string, c context.Context) (error, string)
	Login(req dto.LoginRequest, ctx context.Context) (string, string, error)
	LogOut(id string, ctx context.Context) error
	//VerifyAction(rawToken string, c context.Context) (error, string)
	//ResetPassword(newPass, re_newPass, token string, c context.Context) response.ResetPasswordResponse
	//RecoverAccountByCustomer(email string, c context.Context) (string, error)
}
