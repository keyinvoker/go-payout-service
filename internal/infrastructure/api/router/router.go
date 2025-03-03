package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers"
	v1 "github.com/keyinvoker/go-payout-service/internal/infrastructure/api/handlers/v1"
)

type Router struct {
	engine        *gin.Engine
	healthHandler *handlers.HealthHandler
	payoutHandler *v1.PayoutHandler
}

func NewRouter(
	healthHandler *handlers.HealthHandler,
	payoutHandler *v1.PayoutHandler,
) *Router {
	return &Router{
		engine:        gin.Default(),
		healthHandler: healthHandler,
		payoutHandler: payoutHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {

	api := r.engine.Group("/api")

	{
		api.GET("/healthz", r.healthHandler.CheckHealth)

		apiV1 := api.Group("/v1")
		{
			apiV1.GET("/payout/:id", r.payoutHandler.GetPayoutByID)
			apiV1.POST("/payout", r.payoutHandler.CreatePayout)
			apiV1.PUT("/payout", r.payoutHandler.UpdatePayoutDescription)
		}
	}

	return r.engine
}
