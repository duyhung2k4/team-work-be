package repository

import (
	"errors"
	"team-work-be/config"
	"team-work-be/model"

	"gorm.io/gorm"
)

type accessRepo struct {
	db *gorm.DB
}

type AccessRepo interface {
	CheckUser(username string) (*model.Profile, error)
	GetProfile(userId uint) (*model.Profile, error)
}

func (a *accessRepo) CheckUser(username string) (*model.Profile, error) {
	var profile *model.Profile
	var credential *model.Credential

	if err := a.db.
		Model(&model.Credential{}).
		Where("username = ?", username).
		First(&credential).Error; err != nil {
		return nil, err
	}

	if credential == nil {
		return nil, errors.New("user not found")
	}

	if err := a.db.Model(&model.Profile{}).
		Preload("Credential").
		Preload("Credential.Role").
		Where("credential_id = ?", credential.Id).
		First(&profile).Error; err != nil {
		return nil, err
	}

	if profile == nil {
		return nil, errors.New("user not found")
	}

	return profile, nil
}

func (a *accessRepo) GetProfile(userId uint) (*model.Profile, error) {
	var profile *model.Profile

	if err := a.db.
		Model(&model.Profile{}).
		Preload("Credential").
		Preload("Credential.Role").
		Where("credential_id = ?", userId).
		First(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func NewAccessRepo() AccessRepo {
	db := config.GetDB()
	return &accessRepo{
		db: db,
	}
}
