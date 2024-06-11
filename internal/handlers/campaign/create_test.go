package campaign

import (
	"bytes"
	"emailn/internal/dto"
	"emailn/internal/validator"
	internalMock "emailn/test/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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

func Test_Create(t *testing.T) {
	assert := assert.New(t)

	t.Run("Given valid request When calling Create without errors Then return 201 and id in body", func(t *testing.T) {
		service := new(internalMock.CampaignServiceMock)
		service.On("Create", newCampaign).Return("1", nil)
		handler := &CampaignHandler{
			service: service,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(newCampaign)
		req, _ := http.NewRequest("POST", "/campaigns", &buf)
		res := httptest.NewRecorder()

		responseBody, statusCode, _ := handler.Create(res, req)

		assert.Equal(http.StatusCreated, statusCode)
		assert.Equal(map[string]string{"id": "1"}, responseBody)
	})

	//TODO: Validate domain errors in service return (if is a domain error then return 400 and error in body)
	t.Run("Given any request When calling Create with service error Then return 500", func(t *testing.T) {
		expectedError := errors.New("error")
		service := new(internalMock.CampaignServiceMock)
		service.On("Create", newCampaign).Return("", expectedError)
		handler := &CampaignHandler{
			service: service,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(newCampaign)
		req, _ := http.NewRequest("POST", "/campaigns", &buf)
		res := httptest.NewRecorder()

		responseBody, statusCode, err := handler.Create(res, req)

		assert.Error(err)
		assert.True(errors.Is(err, expectedError))
		assert.Equal(http.StatusInternalServerError, statusCode)
		assert.Nil(responseBody)
	})

	t.Run("Given request with empty body When calling Create Then return 400 and validation error", func(t *testing.T) {
		handler := &CampaignHandler{
			service: new(internalMock.CampaignServiceMock),
		}
		req, _ := http.NewRequest("POST", "/campaigns", nil)
		res := httptest.NewRecorder()

		responseBody, statusCode, err := handler.Create(res, req)

		assert.Error(err)
		assert.Equal(err.Error(), validator.EmptyBodyValidationMessage)
		assert.Equal(http.StatusBadRequest, statusCode)
		assert.Nil(responseBody)
	})

	t.Run("Given request with invalid body field type When calling Create Then return 400 and validation error", func(t *testing.T) {
		handler := &CampaignHandler{
			service: new(internalMock.CampaignServiceMock),
		}
		req, err := http.NewRequest("POST", "/campaigns", strings.NewReader(`{"name": 1, "content": "Content Body", "contacts": ["email1@eee.com", "email2@eee.com"]}`))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		response, statusCode, err := handler.Create(w, req)

		assert.Equal(http.StatusBadRequest, statusCode)
		assert.Contains(err.Error(), "name: Invalid type")
		assert.Nil(response)
	})
}
