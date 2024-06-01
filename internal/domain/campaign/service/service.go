package campaign_service

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
)

type Service interface {
	Create(newCampaign *dto.NewCampaign) (string, error)
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