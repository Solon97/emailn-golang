package campaign

import (
	"emailn/internal/dto"
	"errors"
	"net/http"
)

func (handler *CampaignHandler) Create(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	newCampaign, validationErrorMessage, err := dto.NewCampaignDto(r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if validationErrorMessage != "" {
		return nil, http.StatusBadRequest, errors.New(validationErrorMessage)
	}

	id, err := handler.service.Create(newCampaign)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"id": id}, http.StatusCreated, nil
}
