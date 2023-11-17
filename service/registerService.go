package service

import (
	"errors"
	"team-work-be/model"
	"team-work-be/payload"
	"team-work-be/repository"
	"time"
)

type registerService struct {
	registerRepo repository.RegisterRepo
}

type RegisterService interface {
	CreateTemporaryCredential(infoRegisterPayload payload.InfoRegisterPayload) (*model.TemporaryCredential, error)
	AuthenCodeRegister(codeRegisterPayload payload.CodeRegisterPayload) error
}

func (re *registerService) CreateTemporaryCredential(infoRegisterPayload payload.InfoRegisterPayload) (*model.TemporaryCredential, error) {
	existEmail := re.registerRepo.CheckExistEmail(infoRegisterPayload)

	if existEmail {
		return nil, errors.New("email existed")
	}

	temporaryCredential, err := re.registerRepo.CreateTemporaryCredential(infoRegisterPayload)

	return temporaryCredential, err
}

func (re *registerService) AuthenCodeRegister(codeRegisterPayload payload.CodeRegisterPayload) error {
	temporaryCredential, errCode := re.registerRepo.GetTemporaryCredential(codeRegisterPayload.IdTemporaryCredential)

	if errCode != nil {
		return errCode
	}

	if temporaryCredential.Code != codeRegisterPayload.Code {
		return errors.New("code error")
	}

	if time.Now().After(temporaryCredential.TimeEnd) {
		return errors.New("time error")
	}

	temporaryCredential, errGet := re.registerRepo.GetTemporaryCredential(codeRegisterPayload.IdTemporaryCredential)
	if errGet != nil {
		return errGet
	}

	err := re.registerRepo.CreateCredential(*temporaryCredential)

	return err
}

func NewRegisterService() RegisterService {
	registerRepo := repository.NewRegisterRepo()
	return &registerService{
		registerRepo: registerRepo,
	}
}
