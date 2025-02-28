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

func (h *PayoutHandler) GetPayoutByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payout ID"})
		return
	}

	payout, err := h.payoutService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payout not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payout": payout})
}
