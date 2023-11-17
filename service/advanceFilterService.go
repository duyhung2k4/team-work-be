package service

import (
	"team-work-be/model"
	"team-work-be/repository"
)

type advanceFilterService struct {
	repo repository.AdvanceFilterRepo
}

type AdvanceFilterService interface {
	AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error)
}

func (a *advanceFilterService) AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error) {
	data, total, err := a.repo.AdvanceFilter(payload)

	return data, total, err
}

func NewAdvanceFilterService() AdvanceFilterService {
	repo := repository.NewAdvanceFilterRepo()
	return &advanceFilterService{
		repo: repo,
	}
}
