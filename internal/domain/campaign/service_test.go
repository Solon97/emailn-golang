package campaign

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = &dto.NewCampaign{
		Name:     "Campaign X",
		Content:  "Content Body",
		Contacts: []string{"email1@eee.com", "email2@eee.com"},
	}
)

func Test_Save_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	repo := &repositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	service := NewService(repo)

	id, err := service.Save(newCampaign)

	assert.NoError(err)
	assert.NotEmpty(id)
}

func Test_Save_Repository(t *testing.T) {
	repo := &repositoryMock{}
	repo.On("Save", mock.MatchedBy(func(c *Campaign) bool {
		return c.Name == newCampaign.Name &&
			c.Content == newCampaign.Content &&
			len(c.Contacts) == len(newCampaign.Contacts)
	})).Return(nil)
	service := NewService(repo)

	service.Save(newCampaign)

	repo.AssertExpectations(t)
}

func Test_Save_ValidationError(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "name", "5")
	repo := &repositoryMock{}
	service := NewService(repo)

	_, err := service.Save(&dto.NewCampaign{})

	assert.EqualError(err, expectedError.Error())
}

func Test_Save_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repo := &repositoryMock{}
	repo.On("Save", mock.Anything).Return(errors.New("error"))
	service := NewService(repo)

	_, err := service.Save(newCampaign)
	assert.True(errors.Is(err, internalerrors.ErrInternalServer))
	repo.AssertExpectations(t)
}
