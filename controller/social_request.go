package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func AcceptRequest(ctx *gin.Context) {
	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var requestId = ctx.Param("requestId")
	var actorId = ctx.Param("actorId")

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.AcceptRequest(requestId, actorId, ctx),
		Context: ctx,
	})
}

func CancelRequest(ctx *gin.Context) {
	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var requestId = ctx.Param("requestId")
	var actorId = ctx.Param("actorId")

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.CancelRequest(requestId, actorId, ctx),
		Context: ctx,
	})
}

func GetRequest(ctx *gin.Context) {
	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var id = ctx.Param("id")
	res, err := service.GetRequest(id, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func GetRequestsToUser(ctx *gin.Context) {
	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var id = ctx.Param("id")
	var requestType = ctx.Param("requestType")

	res, err := service.GetRequestsToUser(id, requestType, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func GetAllRequests(ctx *gin.Context) {
	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	res, err := service.GetAllRequests(ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func ProcessRequest(ctx *gin.Context) {
	var request dto.SocialRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateSocialRequestService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.ProcessRequest(request, ctx),
		Context: ctx,
	})
}
