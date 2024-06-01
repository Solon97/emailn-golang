package main

import (
	handler "emailn/internal/handlers/campaign"
	"emailn/internal/infrastructure/database"
	"net/http"

	service "emailn/internal/domain/campaign/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService, err := service.NewCampaignService(&database.CampaignRepository{})
	if err != nil {
		panic(err)
	}

	campaignHandler := handler.NewCampaignHandler(campaignService)

	r.Post("/campaigns", campaignHandler.Create)
	r.Get("/campaigns/{id}", campaignHandler.GetByID)

	http.ListenAndServe(":3000", r)
}
