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

func Test_UpdateSendStatus(t *testing.T) {
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

	assert := assert.New(t)

	t.Run("Success", func(t *testing.T) {
		newStatus := entity.SendStatusStarted
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(campaign, nil)
		repo.On("UpdateSendStatus",
			mock.MatchedBy(func(id string) bool { return id == campaign.ID }),
			mock.MatchedBy(func(status entity.SendStatus) bool { return status == newStatus })).
			Return(nil)
		service, _ := NewCampaignService(repo)
		err := service.UpdateSendStatus(campaign.ID, newStatus)
		assert.NoError(err)
	})

	t.Run("Not found Campaign", func(t *testing.T) {
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(nil, errors.New("campaign not found"))
		service, _ := NewCampaignService(repo)
		err := service.UpdateSendStatus(campaign.ID, entity.SendStatusStarted)
		assert.EqualError(err, "campaign not found")
	})

	t.Run("Internal Server Error in GetById", func(t *testing.T) {
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(campaign, internalerrors.ErrInternalServer)
		service, _ := NewCampaignService(repo)
		err := service.UpdateSendStatus(campaign.ID, entity.SendStatusStarted)
		assert.True(errors.Is(err, internalerrors.ErrInternalServer))
	})

	t.Run("Internal Server Error in UpdateSendStatus", func(t *testing.T) {
		newStatus := entity.SendStatusStarted
		repo := &repositoryMock{}
		repo.On("GetById", mock.Anything).Return(campaign, nil)
		repo.On("UpdateSendStatus",
			mock.MatchedBy(func(id string) bool { return id == campaign.ID }),
			mock.MatchedBy(func(status entity.SendStatus) bool { return status == newStatus })).
			Return(internalerrors.ErrInternalServer)
		service, _ := NewCampaignService(repo)
		err := service.UpdateSendStatus(campaign.ID, newStatus)
		assert.True(errors.Is(err, internalerrors.ErrInternalServer))
	})
}
