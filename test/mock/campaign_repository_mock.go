package mock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) GetAll() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (r *CampaignRepositoryMock) GetById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*campaign.Campaign), args.Error(1)
	}
	return nil, args.Error(1)
}

func (r *CampaignRepositoryMock) UpdateSendStatus(id string, status campaign.SendStatus) error {
	args := r.Called(id, status)
	return args.Error(0)
}
