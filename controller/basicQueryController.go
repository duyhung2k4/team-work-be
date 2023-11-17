package controller

import (
	"encoding/json"
	"net/http"
	"team-work-be/model"
	"team-work-be/service"

	"github.com/go-chi/render"
)

type basicQueryController struct {
	service service.BasicQueryService
}

type BasicQueryController interface {
	BasicQuery(w http.ResponseWriter, r *http.Request)
}

// @Summary      Basic query
// @Description  Basic query
// @Tags         Basic query
// @Accept       json
// @Produce      json
// @Param        req   body model.BasicQueryPayload true "payload"
// @Success      200  {object}  Response
// @Router       /basic-query/query [post]
func (b *basicQueryController) BasicQuery(w http.ResponseWriter, r *http.Request) {
	var payload model.BasicQueryPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	result, err := b.service.BasicQuery(payload)

	if err != nil {
		serverError(w, r, err)
		return
	}

	meta := Response{
		Data:    result,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, meta)
}

func NewBasicQueryController() BasicQueryController {
	service := service.NewBasicQueryService()
	return &basicQueryController{
		service: service,
	}
}
