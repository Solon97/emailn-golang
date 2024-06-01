package campaign

import (
	"emailn/internal/dto"

	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign *dto.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *serviceMock) GetById(id string) (*dto.GetCampaignResponse, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.GetCampaignResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
