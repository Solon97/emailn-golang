package campaign

import (
	"emailn/internal/handlers"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *CampaignHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/campaigns/"):]
	campaign, err := handler.service.GetById(id)
	if err != nil {
		handlers.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaign)
}
