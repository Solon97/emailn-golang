package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body"
	emails  = []string{
		"email1@eee.com",
		"email2@eee.com",
	}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, emails)
	assert.NoError(err)
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, len(emails), len(campaign.Contacts))
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

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign("", content, emails)

	assert.Nil(campaign)
	assert.EqualError(err, "name is required")
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, "", emails)

	assert.Nil(campaign)
	assert.EqualError(err, "content is required")
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, []string{})

	assert.Nil(campaign)
	assert.EqualError(err, "at least one contact is required")
}
