package controller

import (
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context) {
	var request dto.CreateCommentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.PostComment(request, ctx),
		Context: ctx,
	})
}

func RemoveComment(ctx *gin.Context) {
	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.RemoveComment(ctx.Param("actorId"), ctx.Param("commentId"), ctx),
		Context: ctx,
	})
}

func EditComment(ctx *gin.Context) {
	var request dto.EditCommentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.EditComment(request, ctx),
		Context: ctx,
	})
}

func GetCommentsFromPost(ctx *gin.Context) {
	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		Data1:   service.GetCommentsFromPost(ctx.Param("id"), ctx),
		Context: ctx,
	})
}
