package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Create(campaign *campaign.Campaign) error {
	r.campaigns = append(r.campaigns, *campaign)
	return nil
}

func (r *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	return []campaign.Campaign{}, errors.New("not implemented")
}
