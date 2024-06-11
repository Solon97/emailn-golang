package campaign_service

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
)

type Service interface {
	Create(newCampaign *dto.NewCampaign) (string, error)
	GetById(id string) (*dto.GetCampaignResponse, error)
}

type CampaignService struct {
	repository Repository
}

func NewCampaignService(repository Repository) (*CampaignService, error) {
	if repository == nil {
		return nil, internalerrors.ErrRepositoryNil
	}
	// TODO: Create Domain Custom Errors
	return &CampaignService{
		repository: repository,
	}, nil
}
