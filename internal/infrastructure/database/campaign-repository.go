package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	r.campaigns = append(r.campaigns, *campaign)
	return errors.New("error")
}
