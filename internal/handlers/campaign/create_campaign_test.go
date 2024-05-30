package campaign

import (
	"bytes"
	"emailn/internal/domain/campaign"
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = &dto.NewCampaign{
		Name:     "Campaign X",
		Content:  "Content Body",
		Contacts: []string{"email1@eee.com", "email2@eee.com"},
	}
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign *dto.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *serviceMock) GetAll() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func Test_CreateCampaign_SaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	service := &serviceMock{}
	service.On("Create", newCampaign).Return("1", nil)
	handler := &CampaignHandler{
		service: service,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(newCampaign)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	res := httptest.NewRecorder()

	handler.CreateCampaign(res, req)

	assert.Equal(http.StatusCreated, res.Code)
}

func Test_CreateCampaign_ServiceError(t *testing.T) {
	assert := assert.New(t)
	service := &serviceMock{}
	service.On("Create", newCampaign).Return("", errors.New("error"))
	handler := &CampaignHandler{
		service: service,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(newCampaign)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	res := httptest.NewRecorder()

	handler.CreateCampaign(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func Test_CreateCampaign_InternalServerError(t *testing.T) {
	assert := assert.New(t)
	service := &serviceMock{}
	service.On("Create", newCampaign).Return("", internalerrors.ErrInternalServer)
	handler := &CampaignHandler{
		service: service,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(newCampaign)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	res := httptest.NewRecorder()

	handler.CreateCampaign(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
}

func Test_CreateCampaign_EmptyBody(t *testing.T) {
	assert := assert.New(t)
	handler := &CampaignHandler{
		service: &serviceMock{},
	}
	req, _ := http.NewRequest("POST", "/campaigns", nil)
	res := httptest.NewRecorder()
	handler.CreateCampaign(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "empty body")
}

func Test_CreateCampaign_DecodeJSONError(t *testing.T) {
	assert := assert.New(t)
	handler := &CampaignHandler{
		service: &serviceMock{},
	}
	req, err := http.NewRequest("POST", "/campaigns", strings.NewReader(`{"name": 1, "content": "Content Body", "contacts": ["email1@eee.com", "email2@eee.com"]}`))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler.CreateCampaign(w, req)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Contains(w.Body.String(), "invalid name")
}
