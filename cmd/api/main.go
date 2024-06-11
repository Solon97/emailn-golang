package main

import (
	"emailn/internal/handlers"
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

	db := database.NewDB()
	campaignRepository := database.NewCampaignRepository(db)
	campaignService, err := service.NewCampaignService(campaignRepository)
	if err != nil {
		panic(err)
	}

	campaignHandler := handler.NewCampaignHandler(campaignService)

	r.Post("/campaigns", handlers.HandleError(campaignHandler.Create))
	r.Get("/campaigns/{id}", handlers.HandleError(campaignHandler.GetByID))

	http.ListenAndServe(":3000", r)
}
