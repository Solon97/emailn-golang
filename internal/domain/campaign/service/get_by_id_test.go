package campaign_service

import (
	entity "emailn/internal/domain/campaign"
	internalerrors "emailn/internal/errors"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetById(t *testing.T) {
	assert := assert.New(t)
	campaign := &entity.Campaign{
		ID:      "1",
		Name:    "Campaign X",
		Content: "Content Body",
		Contacts: []entity.Contact{
			{Email: "email1@eee.com"},
			{Email: "email2@eee.com"},
		},
		SendStatus: entity.SendStatusPending,
		CreadetAt:  time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		repo := &repositoryMock{}
		repo.On("GetById", mock.MatchedBy(func(id string) bool { return id == campaign.ID })).Return(campaign, nil)
		service, _ := NewCampaignService(repo)

		campaignResponse, err := service.GetById(campaign.ID)

		assert.NoError(err)
		assert.Equal(campaign.ID, campaignResponse.ID)
		assert.Equal(campaign.Name, campaignResponse.Name)
		assert.Equal(campaign.Content, campaignResponse.Content)
		assert.Equal(string(campaign.SendStatus), campaignResponse.SendStatus)
	})

	t.Run("Repository error", func(t *testing.T) {
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(nil, errors.New("repository error"))
		service, _ := NewCampaignService(repo)
		_, err := service.GetById(campaign.ID)

		assert.True(errors.Is(err, internalerrors.ErrInternalServer))
	})

	t.Run("Not found Campaign", func(t *testing.T) {
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(nil, nil)
		service, _ := NewCampaignService(repo)

		campaignResponse, err := service.GetById(campaign.ID)

		assert.NoError(err)
		assert.Nil(campaignResponse)
	})

}
