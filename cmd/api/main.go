package main

import (
	"log"
	"net/http"

	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/repositories"

	v1 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v1"
	v2 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v2"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/persistence/database/postgres"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	payoutRepo := repositories.NewPayoutRepository(db)
	payoutService := services.NewPayoutService(payoutRepo)
	payoutHandler := v1.NewPayoutHandler(payoutService)

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/payouts/:id", payoutHandler.GetPayoutByID)
		// apiV1.POST("/payouts", payoutHandler.CreatePayout)
	}

	muxRouter := mux.NewRouter()
	apiV2 := muxRouter.PathPrefix("/api/v2").Subrouter()

	payoutHandlerV2 := v2.NewPayoutHandler(payoutService)

	apiV2.HandleFunc("/payouts/:id", payoutHandlerV2.GetPayoutByID).Methods("GET")

	port := "8888"
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed:", err)
	}

}
