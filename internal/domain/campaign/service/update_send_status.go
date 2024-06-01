package campaign_service

import (
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/errors"
	"errors"

	"github.com/sirupsen/logrus"
)

func (s *CampaignService) UpdateSendStatus(id string, status campaign.SendStatus) error {
	_, err := s.repository.GetById(id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternalServer) {
			logrus.WithField("function", "UpdateSendStatus").WithError(err).Error("error retrieving campaign")
			return internalerrors.ErrInternalServer
		}
		return err
	}
	if err := s.repository.UpdateSendStatus(id, status); err != nil {
		logrus.WithField("function", "UpdateSendStatus").WithError(err).Error("error updating campaign")
		return internalerrors.ErrInternalServer
	}
	return nil
}
