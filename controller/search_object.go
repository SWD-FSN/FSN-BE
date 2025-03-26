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

	var nullId = ""
	var keyword = ctx.Param("keyword")

	util.ProcessResponse(dto.APIResponse{
		Data1:   service.GetObjectsByKeyword(nullId, keyword, ctx),
		Context: ctx,
	})
}
