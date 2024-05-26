package campaign

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) (*Service, error) {
	if repository == nil {
		return nil, internalerrors.ErrRepositoryNil
	}
	return &Service{
		repository: repository,
	}, nil
}

func (s *Service) Create(newCampaign *dto.NewCampaign) (string, error) {
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

func (s *Service) GetAll() ([]Campaign, error) {
	return s.repository.GetAll()
}
