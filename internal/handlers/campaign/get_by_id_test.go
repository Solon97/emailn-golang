package campaign

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	internalMock "emailn/test/mock"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetCampaignByID(t *testing.T) {
	campaignResponse := &dto.GetCampaignResponse{
		ID:         "1",
		Name:       "Campaign X",
		Content:    "Content Body",
		SendStatus: string(campaign.SendStatusPending),
	}

	assert := assert.New(t)
	t.Run("Success", func(t *testing.T) {
		service := new(internalMock.CampaignServiceMock)
		service.On("GetById", mock.MatchedBy(func(id string) bool { return id == campaignResponse.ID })).Return(campaignResponse, nil)
		handler := &CampaignHandler{
			service: service,
		}
		req, _ := http.NewRequest("POST", fmt.Sprintf("/campaigns/%s", campaignResponse.ID), nil)
		res := httptest.NewRecorder()

		responseBody, statusCode, err := handler.GetByID(res, req)

		assert.Equal(http.StatusOK, statusCode)
		assert.Nil(err)
		assert.Equal(campaignResponse, responseBody)
		service.AssertExpectations(t)
	})

	t.Run("Not found", func(t *testing.T) {
		service := new(internalMock.CampaignServiceMock)
		service.On("GetById", mock.Anything).Return(nil, errors.New("campaign not found"))
		handler := &CampaignHandler{
			service: service,
		}
		req, _ := http.NewRequest("POST", fmt.Sprintf("/campaigns/%s", campaignResponse.ID), nil)
		res := httptest.NewRecorder()

		responseBody, statusCode, err := handler.GetByID(res, req)

		assert.Equal(http.StatusNotFound, statusCode)
		assert.True(strings.Contains(err.Error(), "campaign not found"))
		assert.Nil(responseBody)
		service.AssertExpectations(t)
	})

	t.Run("Service error", func(t *testing.T) {
		service := new(internalMock.CampaignServiceMock)
		service.On("GetById", mock.Anything).Return(nil, internalerrors.ErrInternalServer)
		handler := &CampaignHandler{
			service: service,
		}
		req, _ := http.NewRequest("POST", fmt.Sprintf("/campaigns/%s", campaignResponse.ID), nil)
		res := httptest.NewRecorder()

		responseBody, statusCode, err := handler.GetByID(res, req)

		assert.Equal(http.StatusInternalServerError, statusCode)
		assert.True(strings.Contains(err.Error(), internalerrors.ErrInternalServer.Error()))
		assert.Nil(responseBody)
		service.AssertExpectations(t)
	})
}
