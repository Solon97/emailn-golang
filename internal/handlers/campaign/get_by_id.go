package campaign

import (
	"emailn/internal/handlers"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *CampaignHandler) GetCampaign(w http.ResponseWriter, r *http.Request) {
	campaigns, err := handler.service.GetAll()
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaigns)
}
