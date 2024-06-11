package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(campaign *campaign.Campaign) error {
	return r.db.Create(campaign).Error
}

func (r *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaign *campaign.Campaign
	tx := r.db.Where("id = ?", id).First(&campaign)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}
	return campaign, nil
}

func (r *CampaignRepository) UpdateSendStatus(id string, status campaign.SendStatus) error {
	return r.db.Model(&campaign.Campaign{}).Where("id = ?", id).Update("send_status", status).Error
}
