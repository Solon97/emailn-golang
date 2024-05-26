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

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, emails)
	assert.NoError(err)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(emails), len(campaign.Contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, emails)
	assert.NoError(err)
	assert.NotEmpty(campaign.ID)
}

func Test_NewCampaign_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, err := NewCampaign(name, content, emails)
	assert.NoError(err)
	assert.Greater(campaign.CreadetAt, now)
}

func Test_NewCampaign_MustValidateMinNameLength(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "name", "5")
	campaign, err := NewCampaign("", content, emails)

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}

func Test_NewCampaign_MustValidateMaxNameLength(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMaxFieldPattern, "name", "24")
	campaign, err := NewCampaign(fake.Lorem().Text(30), content, emails)

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}

func Test_NewCampaign_MustValidateMinContentLength(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "content", "5")

	campaign, err := NewCampaign(name, "", emails)

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}

func Test_NewCampaign_MustValidateMaxContentLength(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMaxFieldPattern, "content", "1024")

	campaign, err := NewCampaign(name, fake.Lorem().Text(1100), emails)

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}

func Test_NewCampaign_MustValidateMinEmails(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "contacts", "1")

	campaign, err := NewCampaign(name, content, []string{})

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrEmailFieldPattern, "email")

	campaign, err := NewCampaign(name, content, []string{
		"email1@eee.com",
		"invalid_email",
	})

	assert.Nil(campaign)
	assert.EqualError(err, expectedError.Error())
}
