package campaign

import (
	"net/http"
)

func (handler *CampaignHandler) GetByID(w http.ResponseWriter, r *http.Request) (body any, status int, err error) {
	id := r.URL.Path[len("/campaigns/"):]
	campaign, err := handler.service.GetById(id)
	if err != nil {
		if err.Error() == "campaign not found" {
			return nil, http.StatusNotFound, err
		}
		//TODO: Talvez criar função, para retornar o status de acordo com o tipo do erro. (Ex.: Se repository, 500, se not found, 404, se validation, 400)
		return nil, http.StatusInternalServerError, err
	}
	return campaign, http.StatusOK, nil
}
