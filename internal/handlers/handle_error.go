package handlers

import (
	internalerrors "emailn/internal/errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandleError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, statusCode, err := endpointFunc(w, r)
		if err != nil && !errors.Is(err, internalerrors.ErrInternalServer) {
			responseBody = map[string]string{"error": err.Error()}
		}
		if err != nil && errors.Is(err, internalerrors.ErrInternalServer) {
			logrus.WithError(err).Error("internal server error")
			responseBody = map[string]string{"error": internalerrors.ErrInternalServer.Error()}
		}
		render.Status(r, statusCode)
		if responseBody != nil {
			render.JSON(w, r, responseBody)
		}
	})
}
