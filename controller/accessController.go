package controller

import (
	"encoding/json"
	"net/http"
	"team-work-be/config"
	"team-work-be/service"
	"team-work-be/utils"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type accessController struct {
	service service.AccessService
}

type AccessController interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginWithToken(w http.ResponseWriter, r *http.Request)
}

// @Summary      Login
// @Description  Login
// @Tags         Access
// @Accept       json
// @Produce      json
// @Param        req   body LoginPayload true "payload"
// @Success      200  {object}  Response
// @Router       /access/login [post]
func (a *accessController) Login(w http.ResponseWriter, r *http.Request) {

	var payload LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	profile, errCheckUser := a.service.CheckUser(payload.Username, payload.Password)
	if errCheckUser != nil {
		serverError(w, r, errCheckUser)
		return
	}

	_, accessToken, _ := config.GetJWT().Encode(map[string]interface{}{
		"user_id":  profile.Credential.Id,
		"role":     profile.Credential.Role.Name,
		"username": profile.Credential.Username,
		"uuid":     uuid.New().String(),
	})

	_, refreshToken, _ := config.GetJWT().Encode(map[string]interface{}{
		"user_id":  profile.Credential.Id,
		"role":     profile.Credential.Role.Name,
		"username": profile.Credential.Username,
		"uuid":     uuid.New().String(),
	})

	res := Response{
		Data: map[string]interface{}{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"profile":      profile,
		},
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

// @Summary      Login
// @Description  Login
// @Tags         Access
// @Accept       json
// @Produce      json
// @Param        req   body LoginPayloadToken true "payload"
// @Success      200  {object}  Response
// @Router       /access-protected/login-token [post]
func (a *accessController) LoginWithToken(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetFieldToken(utils.USER_ID, r)
	if err != nil {
		serverError(w, r, err)
	}

	userIdConvertToUint := userId.(float64)
	profile, err := a.service.GetProfile(uint(userIdConvertToUint))

	if err != nil {
		serverError(w, r, err)
	}

	_, accessToken, _ := config.GetJWT().Encode(map[string]interface{}{
		"user_id":  profile.Credential.Id,
		"role":     profile.Credential.Role.Name,
		"username": profile.Credential.Username,
		"uuid":     uuid.New().String(),
	})

	_, refreshToken, _ := config.GetJWT().Encode(map[string]interface{}{
		"user_id":  profile.Credential.Id,
		"role":     profile.Credential.Role.Name,
		"username": profile.Credential.Username,
		"uuid":     uuid.New().String(),
	})

	res := Response{
		Data: map[string]interface{}{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"profile":      profile,
		},
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

func NewAccessController() AccessController {
	service := service.NewAccessService()
	return &accessController{
		service: service,
	}
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPayloadToken struct {
	RefreshToken string `json:"refreshToken"`
}
