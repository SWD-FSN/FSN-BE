package controller

import (
	business_object "social_network/business_object"
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"
	"time"

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

var posts = &[]dto.PostResponse{
	{
		PostId:        "1",
		Content:       "Exploring Golang's powerful features!",
		IsPrivate:     false,
		IsHidden:      false,
		LikeAmount:    120,
		CreatedAt:     time.Now(),
		AuthorId:      "101",
		Username:      "golang_dev",
		ProfileAvatar: "https://example.com/avatar1.png",
	},
	{
		PostId:        "2",
		Content:       "Learning concurrency in Go!",
		IsPrivate:     false,
		IsHidden:      false,
		LikeAmount:    200,
		CreatedAt:     time.Now(),
		AuthorId:      "102",
		Username:      "coder123",
		ProfileAvatar: "https://example.com/avatar2.png",
	},
	{
		PostId:        "3",
		Content:       "Private post example",
		IsPrivate:     true,
		IsHidden:      false,
		LikeAmount:    30,
		CreatedAt:     time.Now(),
		AuthorId:      "103",
		Username:      "hidden_user",
		ProfileAvatar: "https://example.com/avatar3.png",
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

func GetPostsDisplayUI(ctx *gin.Context) {
	// service, err := service.GeneratePostService()

	// if err != nil {
	// 	util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
	// 	return
	// }

	util.ProcessLoginResponse(dto.APIReponse{
		Data1:    posts,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}
