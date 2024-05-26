package campaign

import "emailn/internal/domain/campaign"

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{
		service: service,
	}
}
