package controller

import (
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func GetObjectsByKeyword(ctx *gin.Context) {
	service, err := service.GenerateSearchObjectsService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		Data1:   service.GetObjectsByKeyword(ctx.Param("id"), ctx.Param("keyword"), ctx),
		Context: ctx,
	})
}
