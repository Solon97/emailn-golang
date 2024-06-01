package mock

import (
	"emailn/internal/dto"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaign *dto.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *CampaignServiceMock) GetById(id string) (*dto.GetCampaignResponse, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.GetCampaignResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
