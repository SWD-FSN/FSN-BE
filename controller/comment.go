package controller

import (
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"
	"time"

	"github.com/gin-gonic/gin"
)

var comments = &[]dto.CommentDataResponse{
	{
		AuthorId:      "1",
		Author_avatar: "avatar1.png",
		Username:      "user1",
		CommentId:     "cmt1",
		Content:       "This is the first comment",
		CreatedAt:     time.Now(),
		LikeAmount:    5,
	},
	{
		AuthorId:      "2",
		Author_avatar: "avatar2.png",
		Username:      "user2",
		CommentId:     "cmt2",
		Content:       "This is the second comment",
		CreatedAt:     time.Now(),
		LikeAmount:    10,
	},
	{
		AuthorId:      "3",
		Author_avatar: "avatar3.png",
		Username:      "user3",
		CommentId:     "cmt3",
		Content:       "This is the third comment",
		CreatedAt:     time.Now(),
		LikeAmount:    3,
	},
	{
		AuthorId:      "4",
		Author_avatar: "avatar4.png",
		Username:      "user4",
		CommentId:     "cmt4",
		Content:       "This is the fourth comment",
		CreatedAt:     time.Now(),
		LikeAmount:    7,
	},
	{
		AuthorId:      "5",
		Author_avatar: "avatar5.png",
		Username:      "user5",
		CommentId:     "cmt5",
		Content:       "This is the fifth comment",
		CreatedAt:     time.Now(),
		LikeAmount:    15,
	},
}

func PostComment(ctx *gin.Context) {
	var request dto.CreateCommentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.PostComment(request, ctx),
		Context: ctx,
	})
}

func RemoveComment(ctx *gin.Context) {
	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.RemoveComment(ctx.Param("actorId"), ctx.Param("commentId"), ctx),
		Context: ctx,
	})
}

func EditComment(ctx *gin.Context) {
	var request dto.EditCommentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateCommentService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIReponse{
		ErrMsg:  service.EditComment(request, ctx),
		Context: ctx,
	})
}

func GetCommentsFromPost(ctx *gin.Context) {
	// service, err := service.GenerateCommentService()
	// if err != nil {
	// 	util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
	// 	return
	// }

	// util.ProcessResponse(dto.APIReponse{
	// 	Data1:   service.GetCommentsFromPost(ctx.Param("id"), ctx),
	// 	Context: ctx,
	// })

	util.ProcessResponse(dto.APIReponse{
		Data1:    comments,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}
