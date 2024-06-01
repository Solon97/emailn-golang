package handlers

import (
	internalerrors "emailn/internal/errors"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, internalerrors.ErrInternalServer) {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	if strings.Contains(strings.ToLower(err.Error()), "not found") {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, map[string]string{"error": err.Error()})
}
