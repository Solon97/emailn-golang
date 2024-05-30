package handlers

import (
	internalerrors "emailn/internal/errors"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleError_NoError(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	HandleError(w, r, nil)
	assert.Equal(http.StatusOK, w.Code)
}

func Test_HandleError_InternalServerError(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	HandleError(w, r, internalerrors.ErrInternalServer)

	assert.Equal(http.StatusInternalServerError, w.Code)
	assert.Contains(w.Body.String(), "internal server error")
}

func Test_HandleError_BadRequest(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.New("some error")

	HandleError(w, r, err)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Contains(w.Body.String(), "some error")
}
