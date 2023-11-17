package controller

import (
	"encoding/json"
	"net/http"
	"team-work-be/model"
	"team-work-be/service"

	"github.com/go-chi/render"
)

type advanceFilterController struct {
	service service.AdvanceFilterService
}

type AdvanceFilterController interface {
	AdvanceFilter(w http.ResponseWriter, r *http.Request)
}

// @Summary      Advance Filter
// @Description  Advance Filter
// @Tags         Advance Filter
// @Accept       json
// @Produce      json
// @Param        req   body model.AdvanceFilterPayload true "payload"
// @Success      200  {object}  Response
// @Router       /advance-filter/filter [post]
func (a *advanceFilterController) AdvanceFilter(w http.ResponseWriter, r *http.Request) {
	var payload model.AdvanceFilterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
	}

	data, total, err := a.service.AdvanceFilter(payload)

	if err != nil {
		serverError(w, r, err)
		return
	}

	res := MetaResonse{
		Data:    data,
		Message: "OK",
		Success: true,
		Error:   "",
		Page: Page{
			Page:     payload.Page,
			PageSize: payload.PageSize,
			Total:    uint(total),
		},
	}

	render.JSON(w, r, res)
}

func NewAdvanceFilterController() AdvanceFilterController {
	service := service.NewAdvanceFilterService()
	return &advanceFilterController{
		service: service,
	}
}
