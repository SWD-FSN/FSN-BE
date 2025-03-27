package controller

import (
	actiontype "social_network/constant/action_type"
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

	var userId = ctx.Param("userId")
	var keyword = ctx.Param("keyword")

	util.ProcessResponse(dto.APIResponse{
		Data1:    service.GetObjectsByKeyword(userId, keyword, ctx),
		Context:  ctx,
		PostType: actiontype.Non_post,
	})
}
