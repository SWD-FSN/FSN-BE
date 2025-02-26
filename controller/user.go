package controller

import (
	business_object "social_network/business_object"
	action_type "social_network/constant/action_type"
	"social_network/dto"
	"social_network/service"
	"social_network/util"

	"github.com/gin-gonic/gin"
)

var status bool = false

var sampleUsers = &[]business_object.User{
	{
		UserId:        "1",
		RoleId:        "1",
		FullName:      "Full Name 1",
		Username:      "username 1",
		Email:         "email 1",
		Password:      "password 1",
		ProfileAvatar: "image 1",
		Bio:           "bio 1",
		Friends: &[]string{
			"Friend 1",
			"Friend 2",
			"Friend 3",
		},
		Followers: &[]string{
			"Follower 1",
			"Follower 2",
			"Follower 3",
		},
		Followings: &[]string{
			"Following 1",
			"Following 2",
			"Following 3",
		},
		BlockUsers: &[]string{
			"Block user 1",
			"Block user 2",
			"Block user 3",
		},
		IsPrivate:   &status,
		IsActive:    &status,
		IsActivated: true,
	},

	{
		UserId:        "2",
		RoleId:        "2",
		FullName:      "Full Name 2",
		Username:      "username 2",
		Email:         "email 2",
		Password:      "password 2",
		ProfileAvatar: "image 2",
		Bio:           "bio 2",
		Friends: &[]string{
			"Friend 1",
			"Friend 2",
			"Friend 3",
		},
		Followers: &[]string{
			"Follower 1",
			"Follower 2",
			"Follower 3",
		},
		Followings: &[]string{
			"Following 1",
			"Following 2",
			"Following 3",
		},
		BlockUsers: &[]string{
			"Block user 1",
			"Block user 2",
			"Block user 3",
		},
		IsPrivate:   &status,
		IsActive:    &status,
		IsActivated: true,
	},

	{
		UserId:        "3",
		RoleId:        "3",
		FullName:      "Full Name 3",
		Username:      "username 3",
		Email:         "email 3",
		Password:      "password 3",
		ProfileAvatar: "image 3",
		Bio:           "bio 3",
		Friends: &[]string{
			"Friend 1",
			"Friend 2",
			"Friend 3",
		},
		Followers: &[]string{
			"Follower 1",
			"Follower 2",
			"Follower 3",
		},
		Followings: &[]string{
			"Following 1",
			"Following 2",
			"Following 3",
		},
		BlockUsers: &[]string{
			"Block user 1",
			"Block user 2",
			"Block user 3",
		},
		IsPrivate:   &status,
		IsActive:    &status,
		IsActivated: true,
	},

	{
		UserId:        "4",
		RoleId:        "4",
		FullName:      "Full Name 4",
		Username:      "username 4",
		Email:         "email 4",
		Password:      "password 4",
		ProfileAvatar: "image 4",
		Bio:           "bio 4",
		Friends: &[]string{
			"Friend 1",
			"Friend 2",
			"Friend 3",
		},
		Followers: &[]string{
			"Follower 1",
			"Follower 2",
			"Follower 3",
		},
		Followings: &[]string{
			"Following 1",
			"Following 2",
			"Following 3",
		},
		BlockUsers: &[]string{
			"Block user 1",
			"Block user 2",
			"Block user 3",
		},
		IsPrivate:   &status,
		IsActive:    &status,
		IsActivated: true,
	},

	{
		UserId:        "5",
		RoleId:        "5",
		FullName:      "Full Name 5",
		Username:      "username 5",
		Email:         "email 5",
		Password:      "password 5",
		ProfileAvatar: "image 5",
		Bio:           "bio 5",
		Friends: &[]string{
			"Friend 1",
			"Friend 2",
			"Friend 3",
		},
		Followers: &[]string{
			"Follower 1",
			"Follower 2",
			"Follower 3",
		},
		Followings: &[]string{
			"Following 1",
			"Following 2",
			"Following 3",
		},
		BlockUsers: &[]string{
			"Block user 1",
			"Block user 2",
			"Block user 3",
		},
		IsPrivate:   &status,
		IsActive:    &status,
		IsActivated: true,
	},
}

func GetAllUsers(ctx *gin.Context) {
	service, err := service.GenerateUserService()

	if err != nil {
		util.ProcessResponse(util.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetAllUsers(ctx)

	util.ProcessResponse(dto.APIReponse{
		Data1:    res,
		ErrMsg:   err,
		PostType: action_type.Non_post,
		Context:  ctx,
	})

	// util.ProcessResponse(dto.APIReponse{
	// 	Data1:    sampleUsers,
	// 	PostType: action_type.Non_post,
	// 	Context:  ctx,
	// })
}
