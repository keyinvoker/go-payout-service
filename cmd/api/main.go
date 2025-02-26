package main

import (
	"log"
	"net/http"

	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/repositories"

	v1 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v1"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/persistence/database/postgres"

	"github.com/gin-gonic/gin" // Or your preferred web framework
	// ... other imports for handlers, etc.
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

	api := router.Group("/api")
	{
		api.GET("/payouts/:id", payoutHandler.GetPayoutByID)
		// api.POST("/payouts", payoutHandler.CreatePayout)
	}

	port := "8888"
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed:", err)
	}

}
