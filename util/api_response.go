package util

import (
	"errors"
	"fmt"
	"net/http"
	action_type "social_network/constant/action_type"
	"social_network/constant/noti"
	"social_network/dto"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProcessResponse(data dto.APIReponse) {
	if data.ErrMsg != nil {
		processFailResponse(data.ErrMsg, data.Context)
		return
	}

	if data.PostType != action_type.Non_post {
		processSuccessPostReponse(data.Data2, data.PostType, data.Context)
		return
	}

	processSuccessResponse(data.Data1, data.Context)
}

func GenerateInvalidRequestAndSystemProblemModel(c *gin.Context, err error) dto.APIReponse {
	var errMsg error = err
	if errMsg == nil {
		errMsg = errors.New(noti.GenericsErrorWarnMsg)
	}

	return dto.APIReponse{
		ErrMsg:   errMsg,
		Context:  c,
		PostType: action_type.Non_post,
	}
}

func ProcessLoginResponse(data dto.APIReponse) {
	if data.ErrMsg != nil {
		processFailResponse(data.ErrMsg, data.Context)
		return
	}

	var stringRes1 string = fmt.Sprint(data.Data1)
	var stringRes2 string = fmt.Sprint(data.Data2)

	switch stringRes1 {
	case action_type.ActivateCase:
		data.Context.IndentedJSON(http.StatusContinue, gin.H{"message": stringRes2})
	case action_type.Redirect_post:
		processRedirectResponse(stringRes2, data.Context)
	default:
		data.Context.IndentedJSON(http.StatusOK, gin.H{
			"access_token":  stringRes1,
			"refresh_token": stringRes2,
		})
	}
}

func processFailResponse(err error, c *gin.Context) {
	var errCode int

	switch err.Error() {
	case noti.InternalErr:
		errCode = http.StatusInternalServerError
	case noti.GenericsRightAccessWarnMsg:
		errCode = http.StatusForbidden
	default:
		errCode = http.StatusBadRequest
	}

	if isErrorTypeOfUndefined(err) {
		errCode = http.StatusNotFound
	}

	c.IndentedJSON(errCode, gin.H{"message": err.Error()})
}

func processSuccessPostReponse(res interface{}, postType string, c *gin.Context) {
	switch postType {
	case action_type.Redirect_post:
		processRedirectResponse(fmt.Sprint(res), c)
	case action_type.Inform_post:
		processInformResponse(res, c)
	default:
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func processRedirectResponse(redirectUrl string, c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

func processInformResponse(message interface{}, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprint(message)})
}

func processSuccessResponse(data interface{}, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func isErrorTypeOfUndefined(err error) bool {
	return strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "undefined")
}
