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

func (r *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	return nil, errors.New("not implemented")
}

func (r *CampaignRepository) UpdateSendStatus(id string, status campaign.SendStatus) error {
	return errors.New("not implemented")
}
