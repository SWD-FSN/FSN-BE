package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeSearchObjectAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "search-object"

	var norGroup = server.Group(contextPath)
	norGroup.GET("user/:userId/keyword/:keyword", controller.GetObjectsByKeyword)
}
