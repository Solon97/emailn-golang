package campaign_service

import "emailn/internal/domain/campaign"

type Repository interface {
	Create(campaign *campaign.Campaign) error
	GetById(id string) (*campaign.Campaign, error)
	UpdateSendStatus(id string, status campaign.SendStatus) error
}
