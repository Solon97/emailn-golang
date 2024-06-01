package main

import (
	"emailn/internal/domain/campaign"
	handler "emailn/internal/handlers/campaign"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService, err := campaign.NewCampaignService(&database.CampaignRepository{})
	if err != nil {
		panic(err)
	}

	campaignHandler := handler.NewCampaignHandler(campaignService)

	r.Post("/campaigns", campaignHandler.Create)
	r.Get("/campaigns", campaignHandler.GetCampaign)

	http.ListenAndServe(":3000", r)
}
