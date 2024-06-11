package handlers

import (
	internalerrors "emailn/internal/errors"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleError(t *testing.T) {
	assert := assert.New(t)

	t.Run("Given status and empty error and body When calling HandleError Then return the status", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, http.StatusOK, nil
		}
		handler := HandleError(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusOK, w.Code)
		assert.Empty(w.Body.String())
	})

	t.Run("Given body and status and no error When calling HandleError Then return the body and status", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return map[string]string{"key": "value"}, http.StatusOK, nil
		}
		handler := HandleError(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusOK, w.Code)
		assert.Contains(w.Body.String(), "key")
	})

	t.Run("Given status and generic error When calling HandleError Then return the status and the error message", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, http.StatusBadRequest, errors.New("some error")
		}
		handler := HandleError(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Contains(w.Body.String(), "some error")
	})

	t.Run("Given body and status and error When calling HandleError Then return status and the error message in body", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return map[string]string{"key": "value"}, http.StatusBadRequest, errors.New("some error")
		}
		handler := HandleError(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Contains(w.Body.String(), "some error")
		assert.NotContains(w.Body.String(), "key")
	})

	t.Run("Given body and status and wrapped internal server error When calling HandleError Then return status and the internal server error in body", func(t *testing.T) {
		originalErrorMessage := "some error"
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			err := errors.New(originalErrorMessage)
			return nil, http.StatusBadRequest, fmt.Errorf("%w: %s", internalerrors.ErrInternalServer, err.Error())
		}
		handler := HandleError(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Contains(w.Body.String(), "internal server error")
		assert.NotContains(w.Body.String(), originalErrorMessage)
	})
}
