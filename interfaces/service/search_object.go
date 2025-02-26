package service

import "context"

type ISearchObjectService interface {
	GetObjectsByKeyword(id, keyword string, ctx context.Context)
}
