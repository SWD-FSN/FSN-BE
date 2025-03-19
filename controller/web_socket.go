package controller

import (
	"social_network/dto"
	"social_network/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserConnection(ctx *gin.Context) {
	service.RegisterUserConnection(dto.UserConnectionRequest{
		UserId:  ctx.Param("userId"),
		Request: ctx.Request,
		Writer:  ctx.Writer,
	})
}
