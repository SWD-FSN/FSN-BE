package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeConversationAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "conversations"

	var norGroup = server.Group(contextPath)
	norGroup.GET("/keyword", controller.GetConversationsByKeyword)
	norGroup.GET("", controller.GetInternalConversationUIResponse)
	norGroup.POST("/create", controller.CreateConversation)
	norGroup.GET("/user/:id", controller.GetConversationsFromUser)
	norGroup.POST("/message/create", controller.CreateMessage)
}
