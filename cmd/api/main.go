package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repositories "github.com/keyinvoker/go-payout-service/internal/domain/repositories/database"
	"github.com/keyinvoker/go-payout-service/internal/domain/services"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/api/router"

	"github.com/keyinvoker/go-payout-service/internal/config"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers"
	v1 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v1"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/persistence/database/postgres"
)

func main() {
	cfg := config.LoadConfig()

	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize handlers and services
	healthzHandler := handlers.NewHealthHandler(db)

	payoutRepo := repositories.NewPayoutRepository(db)
	payoutService := services.NewPayoutService(payoutRepo)
	payoutHandler := v1.NewPayoutHandler(payoutService)

	router := router.NewRouter(
		healthzHandler,
		payoutHandler,
	)

	handler := router.SetupRoutes()

	port := cfg.ServerPort
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// Start server in a goroutine
	go func() {
		log.Println("Server running on port", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Listening for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
