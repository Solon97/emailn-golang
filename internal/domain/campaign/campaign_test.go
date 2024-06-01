package campaign

import (
	internalerrors "emailn/internal/errors"
	"fmt"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Content Body"
	emails  = []string{
		"email1@eee.com",
		"email2@eee.com",
	}
	fake = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	t.Run("Success", func(t *testing.T) {
		campaign, err := NewCampaign(name, content, emails)

		assert.NoError(err)
		assert.Equal(name, campaign.Name)
		assert.Equal(content, campaign.Content)
		assert.Equal(len(emails), len(campaign.Contacts))
	})

	t.Run("ID is not nil", func(t *testing.T) {
		campaign, err := NewCampaign(name, content, emails)

		assert.NoError(err)
		assert.NotEmpty(campaign.ID)
	})

	t.Run("Created at must be now", func(t *testing.T) {
		campaign, err := NewCampaign(name, content, emails)

		assert.NoError(err)
		assert.Greater(campaign.CreadetAt, time.Now().Add(-time.Minute))
	})

	t.Run("Must validate min name length", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "name", "5")
		campaign, err := NewCampaign("", content, emails)

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Must validate max name length", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrMaxFieldPattern, "name", "24")
		campaign, err := NewCampaign(fake.Lorem().Text(30), content, emails)

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Must validate min content length", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "content", "5")

		campaign, err := NewCampaign(name, "", emails)

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Must validate max content length", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrMaxFieldPattern, "content", "1024")

		campaign, err := NewCampaign(name, fake.Lorem().Text(1100), emails)

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Must validate min emails", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "contacts", "1")

		campaign, err := NewCampaign(name, content, []string{})

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Must validate emails", func(t *testing.T) {
		expectedError := fmt.Errorf(internalerrors.ErrEmailFieldPattern, "email")

		campaign, err := NewCampaign(name, content, []string{
			"email1@eee.com",
			"invalid_email",
		})

		assert.Nil(campaign)
		assert.EqualError(err, expectedError.Error())
	})

	t.Run("Send status should start with pending", func(t *testing.T) {
		campaign, err := NewCampaign(name, content, emails)
		assert.NoError(err)

		assert.Equal(SendStatusPending, campaign.SendStatus)
	})
}
