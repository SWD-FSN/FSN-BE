package controller

import (
	business_object "social_network/business_object"
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func ActivateRole(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var id = ctx.Param("id")

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.ActivateRole(id, ctx),
		Context: ctx,
	})
}

func CreateRole(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var name = ctx.Param("name")

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.CreateRole(name, ctx),
		Context: ctx,
	})
}

func GetAllRoles(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	res, err := service.GetAllRoles(ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func GetRoleById(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var id = ctx.Param("id")
	res, err := service.GetRoleById(id, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func GetRolesByName(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var name = ctx.Param("name")
	res, err := service.GetRolesByName(name, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func GetRolesByStatus(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var status = ctx.Param("status")
	res, err := service.GetRolesByStatus(status, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})
}

func RemoveRole(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	var name = ctx.Param("name")

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.RemoveRole(name, ctx),
		Context: ctx,
	})
}

func UpdateRole(ctx *gin.Context) {
	var role business_object.Role
	if ctx.ShouldBindJSON(&role) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.UpdateRole(role, ctx),
		Context: ctx,
	})
}
