package dto

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Data1    interface{}
	Data2    interface{}
	ErrMsg   error
	Context  *gin.Context
	PostType string
}
