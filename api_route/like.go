package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeLikeAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "likes"

	var norGroup = server.Group(contextPath)
	norGroup.POST("/do-like", controller.DoLike)
	norGroup.POST("/undo-like/:id", controller.UndoLike)
}
