package service

import (
	"team-work-be/model"
	"team-work-be/repository"
)

type basicQueryService struct {
	repo repository.BasicQueryRepo
}

type BasicQueryService interface {
	BasicQuery(payload model.BasicQueryPayload) (interface{}, error)
}

func (b *basicQueryService) BasicQuery(payload model.BasicQueryPayload) (interface{}, error) {
	var result interface{}
	var err error
	switch payload.Option {
	case model.INSERT:
		result, err = b.repo.Insert(payload.Data, payload.ModelType)
	case model.UPDATE:
		result, err = b.repo.Update(payload.Data, payload.ModelType)
	case model.DELETE:
		result, err = b.repo.Delete(payload.Data, payload.ModelType)
	}

	return result, err
}

func NewBasicQueryService() BasicQueryService {
	repo := repository.NewBasicQueryRepo()
	return &basicQueryService{
		repo: repo,
	}
}
