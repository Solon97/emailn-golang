package campaign_service

import (
	entity "emailn/internal/domain/campaign"
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *entity.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) GetAll() ([]entity.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]entity.Campaign), args.Error(1)
}

func (r *repositoryMock) GetById(id string) (*entity.Campaign, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Campaign), args.Error(1)
	}
	return nil, args.Error(1)
}

func (r *repositoryMock) UpdateSendStatus(id string, status entity.SendStatus) error {
	args := r.Called(id, status)
	return args.Error(0)
}

var (
	newCampaign = &dto.NewCampaign{
		Name:     "Campaign X",
		Content:  "Content Body",
		Contacts: []string{"email1@eee.com", "email2@eee.com"},
	}
)

func Test_NewService(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaignService(nil)

	assert.EqualError(err, internalerrors.ErrRepositoryNil.Error())
}
