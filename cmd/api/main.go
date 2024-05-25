package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/dto"
	"emailn/internal/infrastructure/database"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service, err := campaign.NewService(&database.CampaignRepository{})
	if err != nil {
		panic(err)
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var newCampaign dto.NewCampaign
		err := render.DecodeJSON(r.Body, &newCampaign)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := service.Save(&newCampaign)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternalServer) {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, map[string]string{"error": err.Error()})
				return
			}

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)
}
