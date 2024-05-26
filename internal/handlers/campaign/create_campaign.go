package campaign

import (
	"emailn/internal/dto"
	"emailn/internal/handlers"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var newCampaign dto.NewCampaign
	err := render.DecodeJSON(r.Body, &newCampaign)
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	id, err := handler.service.Create(&newCampaign)
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id})
}
