package controller

import (
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

	util.ProcessResponse(dto.APIReponse{
		Data1:   service.GetPersonalProfile(ctx.Param("actorId"), ctx.Param("userId"), ctx),
		Context: ctx,
	})
}
