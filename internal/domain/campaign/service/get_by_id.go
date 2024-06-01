package campaign_service

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	"errors"

	"github.com/sirupsen/logrus"
)

var ErrCampaignNotFound = errors.New("campaign not found")

func (s *CampaignService) GetById(id string) (*dto.GetCampaignResponse, error) {
	campaign, err := s.repository.GetById(id)
	if err != nil {
		logrus.WithField("function", "GetById").WithError(err).Error("error retrieving campaign")
		return nil, internalerrors.ErrInternalServer
	}
	if campaign == nil {
		return nil, ErrCampaignNotFound
	}
	return &dto.GetCampaignResponse{
		ID:         campaign.ID,
		Name:       campaign.Name,
		Content:    campaign.Content,
		SendStatus: string(campaign.SendStatus),
	}, nil
}
