package campaign

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Create(newCampaign *dto.NewCampaign) (string, error)
	GetAll() ([]Campaign, error)
}

type CampaignService struct {
	repository Repository
}

func NewCampaignService(repository Repository) (*CampaignService, error) {
	if repository == nil {
		return nil, internalerrors.ErrRepositoryNil
	}
	return &CampaignService{
		repository: repository,
	}, nil
}

func (s *CampaignService) Create(newCampaign *dto.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)
	if err != nil {
		return "", err
	}
	if err := s.repository.Create(campaign); err != nil {
		logrus.WithError(err).Error("error saving campaign")
		return "", internalerrors.ErrInternalServer
	}
	return campaign.ID, nil
}

func (s *CampaignService) GetAll() ([]Campaign, error) {
	return s.repository.GetAll()
}
