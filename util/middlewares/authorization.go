package middlewares

import (
	"social_network/constant"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

func Authorize(ctx *gin.Context) {
	// Get token from the header
	var token string = ctx.Request.Header.Get("Authorization")

	var unAuthBodyResponse = util.GetUnAuthBodyResponse(ctx)

	if token == "" {
		util.ProcessResponse(unAuthBodyResponse)
		return
	}

	userId, role, _, err := util.ExtractDataFromToken(token, util.GetLogConfig())
	if err != nil {
		util.ProcessResponse(unAuthBodyResponse)
		return
	}

	ctx.Set("userId", userId)
	ctx.Set("role", role)
	ctx.Next()
}

func AdminAuhthorization(ctx *gin.Context) {
	if ctx.GetString("role") != constant.ADMIN_ROLE {
		util.ProcessResponse(util.GetUnAuthBodyResponse(ctx))
		return
	}

	ctx.Next()
}
