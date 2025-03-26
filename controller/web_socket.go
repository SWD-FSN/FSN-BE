package controller

import (
	"social_network/dto"
	"social_network/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserConnection(ctx *gin.Context) {
	var userId = ctx.Param("userId")

	service.RegisterUserConnection(dto.UserConnectionRequest{
		UserId:  userId,
		Request: ctx.Request,
		Writer:  ctx.Writer,
	})
}
