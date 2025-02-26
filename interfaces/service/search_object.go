package service

import (
	"context"
	"social_network/dto"
)

type ISearchObjectService interface {
	GetObjectsByKeyword(id, keyword string, ctx context.Context) *dto.GetObjectsFromEnterSearchBarResponse
}
