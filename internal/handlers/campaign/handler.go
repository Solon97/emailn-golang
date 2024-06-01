package campaign

import (
	service "emailn/internal/domain/campaign/service"
)

type CampaignHandler struct {
	service service.Service
}

func NewCampaignHandler(service service.Service) *CampaignHandler {
	return &CampaignHandler{
		service: service,
	}
}
