package controller

import (
	business_object "social_network/business_object"
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

var samplePosts = &[]business_object.Post{
	{
		PostId:    "1",
		AuthorId:  "1",
		Content:   "1",
		IsPrivate: false,
		IsHidden:  false,
		Status:    true,
	},

	{
		PostId:    "2",
		AuthorId:  "2",
		Content:   "2",
		IsPrivate: false,
		IsHidden:  false,
		Status:    true,
	},

	{
		PostId:    "3",
		AuthorId:  "3",
		Content:   "3",
		IsPrivate: false,
		IsHidden:  false,
		Status:    true,
	},

	{
		PostId:    "4",
		AuthorId:  "4",
		Content:   "4",
		IsPrivate: false,
		IsHidden:  false,
		Status:    true,
	},

	{
		PostId:    "5",
		AuthorId:  "5",
		Content:   "5",
		IsPrivate: false,
		IsHidden:  false,
		Status:    true,
	},
}

func GetAllPosts(ctx *gin.Context) {
	service, err := service.GeneratePostService()

	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetAllPosts(ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})

	// util.ProcessResponse(dto.APIReponse{
	// 	Data1:    samplePosts,
	// 	PostType: action_type.Non_post,
	// 	Context:  ctx,
	// })
}
