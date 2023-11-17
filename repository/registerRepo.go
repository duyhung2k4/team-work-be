package repository

import (
	"errors"
	"math/rand"
	"strconv"
	"team-work-be/config"
	"team-work-be/model"
	"team-work-be/payload"
	"team-work-be/utils"
	"time"

	"gorm.io/gorm"
)

type registerRepo struct {
	db *gorm.DB
}

type RegisterRepo interface {
	CheckExistEmail(infoRegister payload.InfoRegisterPayload) bool
	GetTemporaryCredential(id uint) (*model.TemporaryCredential, error)
	CreateTemporaryCredential(infoRegister payload.InfoRegisterPayload) (*model.TemporaryCredential, error)
	CreateCredential(temporaryCredential model.TemporaryCredential) error
}

func (re *registerRepo) CheckExistEmail(infoRegister payload.InfoRegisterPayload) bool {
	var credential []model.Credential

	if err := re.db.
		Model(&model.Credential{}).
		Where("email = ? OR username = ?", infoRegister.Email, infoRegister.Username).
		Find(&credential).Error; err != nil || len(credential) > 0 {
		return true
	}

	return false
}

func (re *registerRepo) CreateTemporaryCredential(infoRegister payload.InfoRegisterPayload) (*model.TemporaryCredential, error) {
	var temporaryCredential *model.TemporaryCredential

	var code string = ""

	for i := 0; i < 6; i++ {
		numberRandom := rand.Intn(10)
		code += strconv.Itoa(numberRandom)
	}

	temporaryCredential = &model.TemporaryCredential{
		Username:  infoRegister.Username,
		Password:  infoRegister.Password,
		Email:     infoRegister.Email,
		Code:      code,
		TimeStart: time.Now(),
		TimeEnd:   time.Now().Add(time.Second * 30),
	}

	if err := re.db.
		Model(&model.TemporaryCredential{}).
		Create(&temporaryCredential).Error; err != nil {
		return nil, err
	}

	return temporaryCredential, nil
}

func (re *registerRepo) GetTemporaryCredential(id uint) (*model.TemporaryCredential, error) {
	var temporaryCredential *model.TemporaryCredential

	if err := re.db.
		Model(&model.TemporaryCredential{}).
		Where("id = ?", id).
		First(&temporaryCredential).Error; err != nil {
		return nil, err
	}

	return temporaryCredential, nil
}

func (re *registerRepo) CreateCredential(temporaryCredential model.TemporaryCredential) error {
	password, errPassword := utils.HashPassword(temporaryCredential.Password)
	var roleUser *model.Role

	if errPassword != nil {
		return errors.New("error password")
	}

	if errRoleUser := re.db.
		Model(&model.Role{}).
		Where("code = ?", model.USER).
		First(&roleUser).Error; errRoleUser != nil || roleUser == nil {
		return errors.New("error role")
	}

	var credential *model.Credential = &model.Credential{
		Username: temporaryCredential.Username,
		Password: password,
		Email:    temporaryCredential.Email,
		RoleId:   roleUser.Id,
	}

	errCredential := re.db.Model(&model.Credential{}).Create(&credential).Error
	if errCredential != nil {
		return errCredential
	}

	var profile *model.Profile = &model.Profile{
		CredentialId: credential.Id,
	}

	errProfile := re.db.Model(&model.Profile{}).Create(&profile).Error
	if errProfile != nil {
		return errProfile
	}

	return nil
}

func NewRegisterRepo() RegisterRepo {
	db := config.GetDB()
	return &registerRepo{
		db: db,
	}
}
