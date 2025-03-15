package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func DoLike(ctx *gin.Context) {
	var request dto.DoLikeReq
	if ctx.ShouldBindJSON(&request) != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := service.GenerateLikeService()
	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.DoLike(request, ctx),
		Context: ctx,
	})
}

func UndoLike(ctx *gin.Context) {
	service, err := service.GenerateLikeService()
	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	util.ProcessResponse((dto.APIReponse{
		ErrMsg:  service.UndoLike(ctx.Param("id"), ctx),
		Context: ctx,
	}))
}

func GetLikesFromObject(ctx *gin.Context) {
	service, err := service.GenerateLikeService()
	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetLikesFromObject(ctx.Param("id"), ctx.Param("kind"), ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func GetAllLikes(ctx *gin.Context) {
	service, err := service.GenerateLikeService()
	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetAllLikes(ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func GetLike(ctx *gin.Context) {
	service, err := service.GenerateLikeService()
	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetLike(ctx.Param("id"), ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}
