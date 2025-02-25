package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	service, err := service.GenerateUserService()

	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetAllUsers(ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}
