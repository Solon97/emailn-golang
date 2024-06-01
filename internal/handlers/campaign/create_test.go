package campaign

import (
	"bytes"
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
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

	t.Run("Success", func(t *testing.T) {
		service := &serviceMock{}
		service.On("Create", newCampaign).Return("1", nil)
		handler := &CampaignHandler{
			service: service,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(newCampaign)
		req, _ := http.NewRequest("POST", "/campaigns", &buf)
		res := httptest.NewRecorder()

		handler.Create(res, req)

		assert.Equal(http.StatusCreated, res.Code)
	})

	t.Run("Service error", func(t *testing.T) {
		service := &serviceMock{}
		service.On("Create", newCampaign).Return("", errors.New("error"))
		handler := &CampaignHandler{
			service: service,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(newCampaign)
		req, _ := http.NewRequest("POST", "/campaigns", &buf)
		res := httptest.NewRecorder()

		handler.Create(res, req)

		assert.Equal(http.StatusBadRequest, res.Code)
	})

	t.Run("Empty body", func(t *testing.T) {
		handler := &CampaignHandler{
			service: &serviceMock{},
		}
		req, _ := http.NewRequest("POST", "/campaigns", nil)
		res := httptest.NewRecorder()
		handler.Create(res, req)
		assert.Equal(http.StatusBadRequest, res.Code)
		assert.Contains(res.Body.String(), "empty request body")
	})

	t.Run("Internal server error", func(t *testing.T) {
		service := &serviceMock{}
		service.On("Create", newCampaign).Return("", internalerrors.ErrInternalServer)
		handler := &CampaignHandler{
			service: service,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(newCampaign)
		req, _ := http.NewRequest("POST", "/campaigns", &buf)
		res := httptest.NewRecorder()

		handler.Create(res, req)

		assert.Equal(http.StatusInternalServerError, res.Code)
	})

	t.Run("Invalid body field type", func(t *testing.T) {
		handler := &CampaignHandler{
			service: &serviceMock{},
		}
		req, err := http.NewRequest("POST", "/campaigns", strings.NewReader(`{"name": 1, "content": "Content Body", "contacts": ["email1@eee.com", "email2@eee.com"]}`))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handler.Create(w, req)

		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Contains(w.Body.String(), "name: Invalid type")
	})
}
