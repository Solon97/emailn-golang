package campaign

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/internal-errors"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Save(newCampaign *dto.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)
	if err != nil {
		return "", err
	}
	if err := s.repository.Save(campaign); err != nil {
		logrus.WithError(err).Error("error saving campaign")
		return "", internalerrors.ErrInternalServer
	}
	return campaign.ID, nil
}
