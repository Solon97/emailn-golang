package campaign

import (
	"emailn/internal/dto"
	"emailn/internal/handlers"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	newCampaign, validationErrorMessage, err := dto.NewCampaignDto(r.Body)
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	if validationErrorMessage != "" {
		handlers.HandleError(w, r, err)
		return
	}

	id, err := handler.service.Create(newCampaign)
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id})
}
