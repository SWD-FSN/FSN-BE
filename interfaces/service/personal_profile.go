package service

import (
	"context"
	"social_network/dto"
)

type IPersonalProfileService interface {
	GetPersonalProfile(actorId, userId string, ctx context.Context) *dto.PersonalProfileUIResponse
}
