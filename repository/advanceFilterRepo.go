package repository

import (
	"team-work-be/config"
	"team-work-be/model"
	"team-work-be/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type advanceFilterRepo struct {
	db *gorm.DB
}

type AdvanceFilterRepo interface {
	AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error)
}

func (a *advanceFilterRepo) AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error) {
	dataQuery := model.LIST_MODEL_TYPE[payload.ModelType]
	var total int64

	exprs := make([]clause.Expression, 0)

	for key, value := range utils.ConvertNameField(payload.Conditions, model.MODEL_TO_TABLE) {
		exprs = append(exprs, clause.Eq{Column: key, Value: value})
	}

	if len(exprs) == 0 {
		exprs = append(exprs, clause.Eq{Column: "deleted_at", Value: nil})
	}

	conditions := clause.AndConditions{
		Exprs: exprs,
	}

	err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model.MODEL_TYPE[payload.ModelType]).
			Clauses(conditions).
			Count(&total).Error; err != nil {
			return err
		}

		q := tx.Model(model.MODEL_TYPE[payload.ModelType])

		if payload.IsPreload {
			q.Preload(clause.Associations)
			for _, s := range payload.StringPreLoad {
				q.Preload(s)
			}
		}

		if err := q.Clauses(conditions).
			Limit(payload.PageSize).
			Offset((payload.Page - 1) * payload.PageSize).
			Find(&dataQuery).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return dataQuery, total, nil
}

func NewAdvanceFilterRepo() AdvanceFilterRepo {
	db := config.GetDB()
	return &advanceFilterRepo{
		db: db,
	}
}
