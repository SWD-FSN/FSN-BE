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

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.ActivateRole(ctx.Param("id"), ctx),
		Context: ctx,
	})
}

func CreateRole(ctx *gin.Context) {
	service, err := service.GenerateRoleService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.CreateRole(ctx.Param("name"), ctx),
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

	util.ProcessResponse(dto.APIReponse{
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

	res, err := service.GetRoleById(ctx.Param("id"), ctx)

	util.ProcessResponse(dto.APIReponse{
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

	res, err := service.GetRolesByName(ctx.Param("name"), ctx)

	util.ProcessResponse(dto.APIReponse{
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

	res, err := service.GetRolesByStatus(ctx.Param("status"), ctx)

	util.ProcessResponse(dto.APIReponse{
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

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.RemoveRole(ctx.Param("name"), ctx),
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

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.UpdateRole(role, ctx),
		Context: ctx,
	})
}
