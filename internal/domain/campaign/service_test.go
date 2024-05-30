package campaign

import (
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) GetAll() ([]Campaign, error) {
	args := r.Called()
	return args.Get(0).([]Campaign), args.Error(1)
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

func Test_Create_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	repo := &repositoryMock{}
	repo.On("Create", mock.Anything).Return(nil)
	service, _ := NewCampaignService(repo)

	id, err := service.Create(newCampaign)

	assert.NoError(err)
	assert.NotEmpty(id)
}

func Test_Create_UseRepository(t *testing.T) {
	repo := &repositoryMock{}
	repo.On("Create", mock.MatchedBy(func(c *Campaign) bool {
		return c.Name == newCampaign.Name &&
			c.Content == newCampaign.Content &&
			len(c.Contacts) == len(newCampaign.Contacts)
	})).Return(nil)
	service, _ := NewCampaignService(repo)

	service.Create(newCampaign)

	repo.AssertExpectations(t)
}

func Test_Create_ValidationError(t *testing.T) {
	assert := assert.New(t)
	expectedError := fmt.Errorf(internalerrors.ErrMinFieldPattern, "name", "5")
	repo := &repositoryMock{}
	service, _ := NewCampaignService(repo)

	_, err := service.Create(&dto.NewCampaign{})

	assert.EqualError(err, expectedError.Error())
}

func Test_Create_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repo := &repositoryMock{}
	repo.On("Create", mock.Anything).Return(errors.New("error"))
	service, _ := NewCampaignService(repo)

	_, err := service.Create(newCampaign)
	assert.True(errors.Is(err, internalerrors.ErrInternalServer))
	repo.AssertExpectations(t)
}
