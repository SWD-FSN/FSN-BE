package apiroute

import (
	"fmt"
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
	"social_network/controller"
	"social_network/util/middlewares"

	"github.com/gin-gonic/gin"
)

const (
	backUpPort string = "8080"
)

func InitializeUserAPIRoute(server *gin.Engine, logger *log.Logger, port string) {
	if port == "" {
		logger.Println(noti.ApiPortNotSetMsg)

		var apiPortName string = "API port"
		if err := os.Setenv(env.API_PORT, backUpPort); err != nil {
			logger.Println(fmt.Sprintf(noti.EnvSetErrMsg, apiPortName, backUpPort) + err.Error())
		}

		port = backUpPort
	}

	var contextPath string = "users"

	var adminAuthGroup = server.Group(contextPath, middlewares.Authorize, middlewares.AdminAuthorization)
	adminAuthGroup.GET("", controller.GetAllUsers)
	adminAuthGroup.GET("/role/:role", controller.GetUsersByRole)
	adminAuthGroup.GET("/status/:status", controller.GetUsersByStatus)

	var authGroup = server.Group(contextPath, middlewares.Authorize)
	authGroup.GET("/:id", controller.GetUser)
	authGroup.PUT("", controller.UpdateUser)
	//authGroup.PUT("/id/:id/status/:status", controller.ChangeUserStatus)
	authGroup.PUT("/logout", controller.LogOut)

	var norGroup = server.Group(contextPath)
	norGroup.PUT("/login", controller.Login)
	norGroup.POST("/register", controller.CreateUser)
	//norGroup.PUT("/:email", controller.Re)
	norGroup.PUT("/password/:password/confirm-password/:confirmPassword")
	norGroup.PUT("/verify-action", controller.VerifyAction)
}
