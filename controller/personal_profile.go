package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func GetPersonalProfile(ctx *gin.Context) {
	service, err := service.GeneratePersonalProfileService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var actorId = ctx.Param("actorId")
	var userId = ctx.Param("userId")

	util.ProcessResponse(dto.APIResponse{
		Data1:    service.GetPersonalProfile(actorId, userId, ctx),
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}
