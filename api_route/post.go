package apiroute

import (
	"log"
	"social_network/controller"

	"github.com/gin-gonic/gin"
)

func InitializePostAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		port = backUpPort
	}

	var contextPath string = "posts"

	var norGroup = server.Group(contextPath)
	norGroup.GET("", controller.GetPostsDisplayUI)
	norGroup.POST("/create", controller.CreatePost)
	norGroup.PUT("/edit", controller.EditPost)
	norGroup.DELETE("/delete/post/:postId/actor/:actorId", controller.RemovePost)
}
