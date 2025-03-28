package controller

import (
	"fmt"
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"
	"time"

	"github.com/gin-gonic/gin"
)

var conversations = &[]dto.ConversationSearchBarResponse{
	{
		ConversationId:     "1",
		ConversationAvatar: "https://example.com/avatar1.jpg",
		ConversationName:   "John Doe",
	},
	{
		ConversationId:     "2",
		ConversationAvatar: "https://example.com/avatar2.jpg",
		ConversationName:   "Jane Smith",
	},
	{
		ConversationId:     "3",
		ConversationAvatar: "https://example.com/avatar3.jpg",
		ConversationName:   "Alice Brown",
	},
	{
		ConversationId:     "4",
		ConversationAvatar: "https://example.com/avatar4.jpg",
		ConversationName:   "Bob White",
	},
}

func GetConversationsByKeyword(ctx *gin.Context) {
	util.ProcessResponse(dto.APIResponse{
		Data1:    conversations,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func GetInternalConversationUIResponse(ctx *gin.Context) {
	var messages []dto.MessageUIResponse
	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		messages = append(messages, dto.MessageUIResponse{
			MessageId:     fmt.Sprintf("msg_%d", i),
			ConvesationId: "conv_1",
			AuthorId:      fmt.Sprintf("user_%d", i%3+1), // Cycling through 3 users
			AuthorAvatar:  fmt.Sprintf("https://example.com/avatar%d.jpg", i%3+1),
			Content:       fmt.Sprintf("Message number %d", i),
			CreatedAt:     startTime.Add(time.Duration(i) * time.Minute), // Different timestamps
		})
	}

	util.ProcessResponse(dto.APIResponse{
		Data1: dto.InternalConversationUIResponseV2{
			ConversationId:     "conv_1",
			ConversationAvatar: "https://example.com/group_avatar.jpg",
			ConversationName:   "Chat Group",
			RequestUserid:      "user_1",
			Messages:           &messages,
		},
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func CreateConversation(ctx *gin.Context) {
	var request dto.CreateConversationRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateConversationService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	res, err := service.CreateConversation(request, ctx)

	util.ProcessResponse(dto.APIResponse{
		Data1:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}

func CreateMessage(ctx *gin.Context) {
	var request dto.CreateMessageRequest
	if ctx.ShouldBindJSON(&request) != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, nil)
		return
	}

	service, err := service.GenerateConversationService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIResponse{
		ErrMsg:  service.CreateMessage(request, ctx),
		Context: ctx,
	})
}

func GetConversationsFromUser(ctx *gin.Context) {
	service, err := service.GenerateConversationService()
	if err != nil {
		util.GenerateInvalidRequestAndSystemProblemModel(ctx, err)
		return
	}

	util.ProcessResponse(dto.APIResponse{
		Data1:    service.GetConversationsFromUser(ctx.Param("id"), ctx),
		Context:  ctx,
		PostType: action_type.Non_post,
	})
}
