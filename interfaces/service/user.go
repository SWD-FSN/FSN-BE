package service

import (
	"context"
	business_object "social_network/business_object"
	"social_network/dto"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) (*[]business_object.User, error)
	GetUsersByRole(role string, ctx context.Context) (*[]business_object.User, error)
	GetUsersByStatus(rawStatus string, ctx context.Context) (*[]business_object.User, error)
	GetInvoledAccountsFromUser(req dto.GetInvoledAccouuntsRequest, ctx context.Context) (*[]business_object.User, error)
	GetInvolvedAccountsFromTag(id, keyword string, ctx context.Context) *[]dto.GetInvolvedAccountsSearchResponse
	GetUsersFromSearchBar(id, keyword string, ctx context.Context) *[]dto.GetInvolvedAccountsSearchResponse
	GetUser(id string, ctx context.Context) (*business_object.User, error)
	CreateUser(req dto.CreateUserReq, actorId string, ctx context.Context) (string, error)
	UpdateUser(req dto.UpdateUserReq, actorId string, ctx context.Context) (string, error)
	ChangeUserStatus(rawStatus, userId, actorId string, ctx context.Context) (string, error)
	Login(req dto.LoginRequest, ctx *gin.Context) (string, string, error)
	Logout(id string, ctx context.Context) error
	VerifyAction(rawToken string, ctx context.Context) (string, error)
	ResetPassword(newPass, re_newPass, token string, ctx context.Context) (string, error)
	//RecoverAccountByCustomer(email string, c context.Context) (string, error)
}
