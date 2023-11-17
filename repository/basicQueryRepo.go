package repository

import (
	"reflect"
	"team-work-be/config"
	"team-work-be/model"
	"team-work-be/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type basicQueryRepo struct {
	db *gorm.DB
}

type BasicQueryRepo interface {
	Insert(data interface{}, modelType string) (interface{}, error)
	Update(data interface{}, modelType string) (interface{}, error)
	Delete(data interface{}, modelType string) (interface{}, error)
}

func (b *basicQueryRepo) Insert(data interface{}, modelType string) (interface{}, error) {
	if !utils.IsSlice(data) {
		dataInsert := data.(map[string]interface{})
		dataInsert["createdAt"] = time.Now()
		dataInsert["updatedAt"] = time.Now()

		dataInsert = utils.ConvertNameField(dataInsert, model.MODEL_TO_TABLE)

		if err := b.db.Model(model.MODEL_TYPE[modelType]).Clauses(clause.Returning{}).Create(&dataInsert).Error; err != nil {
			return nil, err
		}

		return utils.ConvertNameField(dataInsert, model.TABLE_TO_MODEL), nil
	} else {
		sliceData := reflect.ValueOf(data)
		listInterface := make([]interface{}, sliceData.Len())

		for i := 0; i < sliceData.Len(); i++ {
			listInterface[i] = sliceData.Index(i).Interface()
		}

		listDataInsert := make([]map[string]interface{}, sliceData.Len())

		for i := 0; i < sliceData.Len(); i++ {
			convertMap := listInterface[i].(map[string]interface{})
			convertMap["createdAt"] = time.Now()
			convertMap["updatedAt"] = time.Now()
			listDataInsert[i] = utils.ConvertNameField(convertMap, model.MODEL_TO_TABLE)
		}

		if err := b.db.Model(model.MODEL_TYPE[modelType]).Clauses(clause.Returning{}).Create(&listDataInsert).Error; err != nil {
			return nil, err
		}

		lenData := len(listDataInsert)

		for i := 0; i < lenData; i++ {
			listDataInsert[i] = utils.ConvertNameField(listDataInsert[i], model.TABLE_TO_MODEL)
		}

		return listDataInsert[(lenData / 2):lenData], nil
	}
}

func (b *basicQueryRepo) Update(data interface{}, modelType string) (interface{}, error) {
	if !utils.IsSlice(data) {
		convertMapData := data.(map[string]interface{})
		convertTypeForTable := utils.ConvertNameField(convertMapData, model.MODEL_TO_TABLE)

		if err := b.db.
			Model(model.MODEL_TYPE[modelType]).
			Clauses(
				clause.OnConflict{UpdateAll: true},
				clause.Returning{},
			).
			Create(&convertTypeForTable).Error; err != nil {
			return nil, err
		}

		return utils.ConvertNameField(convertTypeForTable, model.TABLE_TO_MODEL), nil
	} else {
		sliceData := reflect.ValueOf(data)
		listInterface := make([]interface{}, sliceData.Len())
		for i := 0; i < sliceData.Len(); i++ {
			listInterface[i] = sliceData.Index(i).Interface()
		}

		listDataInsert := make([]map[string]interface{}, sliceData.Len())

		for i := 0; i < sliceData.Len(); i++ {
			dataConvertForTable := utils.ConvertNameField(listInterface[i].(map[string]interface{}), model.MODEL_TO_TABLE)
			listDataInsert[i] = dataConvertForTable
		}

		if err := b.db.
			Model(model.MODEL_TYPE[modelType]).
			Clauses(
				clause.OnConflict{UpdateAll: true},
				clause.Returning{},
			).
			Create(&listDataInsert).Error; err != nil {
			return nil, err
		}

		lenData := len(listDataInsert)

		for i := 0; i < lenData; i++ {
			listDataInsert[i] = utils.ConvertNameField(listDataInsert[i], model.TABLE_TO_MODEL)
		}

		return listDataInsert[lenData/2 : lenData], nil
	}
}

func (b *basicQueryRepo) Delete(data interface{}, modelType string) (interface{}, error) {
	if !utils.IsSlice(data) {
		convertMapData := data.(map[string]interface{})
		convertMapData["deletedAt"] = time.Now()

		convertTypeForTable := utils.ConvertNameField(convertMapData, model.MODEL_TO_TABLE)

		if err := b.db.
			Model(model.MODEL_TYPE[modelType]).
			Clauses(
				clause.OnConflict{UpdateAll: true},
				clause.Returning{},
			).
			Create(&convertTypeForTable).Error; err != nil {
			return nil, err
		}

		return utils.ConvertNameField(convertTypeForTable, model.TABLE_TO_MODEL), nil
	} else {
		sliceData := reflect.ValueOf(data)
		listInterface := make([]interface{}, sliceData.Len())
		for i := 0; i < sliceData.Len(); i++ {
			listInterface[i] = sliceData.Index(i).Interface()
		}

		listDataInsert := make([]map[string]interface{}, sliceData.Len())

		for i := 0; i < sliceData.Len(); i++ {
			dataConvertMap := listInterface[i].(map[string]interface{})
			dataConvertMap["deletedAt"] = time.Now()

			dataConvertForTable := utils.ConvertNameField(dataConvertMap, model.MODEL_TO_TABLE)
			listDataInsert[i] = dataConvertForTable
		}

		if err := b.db.
			Model(model.MODEL_TYPE[modelType]).
			Clauses(
				clause.OnConflict{UpdateAll: true},
				clause.Returning{},
			).
			Create(&listDataInsert).Error; err != nil {
			return nil, err
		}

		lenData := len(listDataInsert)

		for i := 0; i < lenData; i++ {
			listDataInsert[i] = utils.ConvertNameField(listDataInsert[i], model.TABLE_TO_MODEL)
		}

		return listDataInsert[lenData/2 : lenData], nil
	}
}

func NewBasicQueryRepo() BasicQueryRepo {
	db := config.GetDB()
	return &basicQueryRepo{
		db: db,
	}
}
