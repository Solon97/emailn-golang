package campaign_service

import (
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/errors"
	"errors"

	"github.com/sirupsen/logrus"
)

var ErrCampaignNotFound = errors.New("campaign not found")

func (s *CampaignService) UpdateSendStatus(id string, status campaign.SendStatus) error {
	campaign, err := s.repository.GetById(id)
	if err != nil {
		return err
	}
	if campaign == nil {
		return ErrCampaignNotFound
	}
	if err := s.repository.UpdateSendStatus(campaign.ID, status); err != nil {
		logrus.WithField("function", "UpdateSendStatus").WithError(err).Error("error updating campaign")
		return internalerrors.ErrInternalServer
	}
	return nil
}
