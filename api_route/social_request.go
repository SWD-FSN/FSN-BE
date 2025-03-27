package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializeSocialRequestAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "social-request"

	var norGroup = server.Group(contextPath)
	norGroup.GET("/user/:id/requestType/:requestType", controller.GetRequestsToUser)
	norGroup.POST("/create", controller.ProcessRequest)
	norGroup.POST("/accept/request/:requestId/actor/:actorId", controller.AcceptRequest)
	norGroup.POST("/cancel/request/:requestId/actor/:actorId", controller.CancelRequest)
}
