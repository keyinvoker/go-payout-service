package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keyinvoker/go-payout-service/internal/application/services"
)

// PayoutHandler defines handlers for Payout operations
type PayoutHandler struct {
	payoutService *services.PayoutService
}

// NewPayoutHandler initializes a new handler
func NewPayoutHandler(payoutService *services.PayoutService) *PayoutHandler {
	return &PayoutHandler{payoutService: payoutService}
}

// GetPayoutByID retrieves a payout by ID
func (handler *PayoutHandler) GetPayoutByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payout ID"})
		return
	}

	payout, err := handler.payoutService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Payout not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payout": payout})
}

// CreatePayout creates a new payout
// func (handler *PayoutHandler) CreatePayout(ctx *gin.Context) {
// 	var payout models.Payout

// 	if err := ctx.ShouldBindJSON(&payout); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	err := handler.payoutService.Create(ctx.Request.Context(), &payout)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payout"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Payout created successfully", "payout": payout})
// }
