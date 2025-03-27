package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeNotificationRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "notifications"

	var norGroup = server.Group(contextPath)
	norGroup.GET("/user/:id", controller.GetUserNotification)
}
