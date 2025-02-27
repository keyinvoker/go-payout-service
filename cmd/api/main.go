package main

import (
	"log"
	"net/http"

	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/repositories"

	v1 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v1"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/persistence/database/postgres"

	"github.com/gorilla/mux"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	payoutRepo := repositories.NewPayoutRepository(db)
	payoutService := services.NewPayoutService(payoutRepo)

	router := mux.NewRouter()
	apiV2 := router.PathPrefix("/api/v1").Subrouter()

	payoutHandlerV1 := v1.NewPayoutHandler(payoutService)

	apiV2.HandleFunc("/payouts/:id", payoutHandlerV1.GetPayoutByID).Methods("GET")

	port := "8888"
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed:", err)
	}

}
