package controller

import (
	"encoding/json"
	"net/http"
	"team-work-be/model"
	"team-work-be/payload"
	"team-work-be/service"
	"team-work-be/utils"

	"github.com/go-chi/render"
)

type registerController struct {
	registerService service.RegisterService
}

type RegisterController interface {
	SendInfoRegister(w http.ResponseWriter, r *http.Request)
	SendCodeRegister(w http.ResponseWriter, r *http.Request)
}

// @Summary      Send Info Register
// @Description  Register
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        req   body payload.InfoRegisterPayload true "payload"
// @Success      200  {object}  Response
// @Router       /access/register/send-info-register [post]
func (re *registerController) SendInfoRegister(w http.ResponseWriter, r *http.Request) {
	var infoRegister payload.InfoRegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&infoRegister); err != nil {
		badRequest(w, r, err)
		return
	}

	temporaryCredential, err := re.registerService.CreateTemporaryCredential(infoRegister)

	if err != nil {
		serverError(w, r, err)
		return
	}

	errSendCode := utils.SendEmail(temporaryCredential.Email, temporaryCredential.Code)

	if errSendCode != nil {
		serverError(w, r, errSendCode)
		return
	}

	responseClient := model.TemporaryCredential{
		Id:        temporaryCredential.Id,
		Username:  temporaryCredential.Username,
		Email:     temporaryCredential.Email,
		TimeStart: temporaryCredential.TimeStart,
		TimeEnd:   temporaryCredential.TimeEnd,
	}

	res := Response{
		Data:    responseClient,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

// @Summary      Send Code Register
// @Description  Register
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        req   body payload.CodeRegisterPayload true "payload"
// @Success      200  {object}  Response
// @Router       /access/register/send-code [post]
func (re *registerController) SendCodeRegister(w http.ResponseWriter, r *http.Request) {
	var codeRegsiter payload.CodeRegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&codeRegsiter); err != nil {
		badRequest(w, r, err)
		return
	}

	errAuthen := re.registerService.AuthenCodeRegister(codeRegsiter)
	if errAuthen != nil {
		serverError(w, r, errAuthen)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func NewRegisterController() RegisterController {
	registerService := service.NewRegisterService()
	return &registerController{
		registerService: registerService,
	}
}
