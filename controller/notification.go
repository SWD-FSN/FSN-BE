package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func GetUserNotification(ctx *gin.Context) {
	service, err := service.GenerateNotiService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	res := service.GetUserNotifications(ctx.Param("id"), ctx)
	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		Data2:    res,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}
