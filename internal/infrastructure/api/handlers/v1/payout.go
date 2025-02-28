package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keyinvoker/go-payout-service/internal/application/services"
)

type PayoutHandler struct {
	payoutService *services.PayoutService
}

func NewPayoutHandler(payoutService *services.PayoutService) *PayoutHandler {
	return &PayoutHandler{payoutService: payoutService}
}

func (h *PayoutHandler) GetPayoutByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payout ID"})
		return
	}

	payout, err := h.payoutService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Payout not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payout": payout})
}
