package campaign_service

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
