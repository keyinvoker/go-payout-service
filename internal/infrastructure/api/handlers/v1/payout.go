package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keyinvoker/go-payout-service/internal/application/dtos"
	"github.com/keyinvoker/go-payout-service/internal/application/usecases"
	"github.com/keyinvoker/go-payout-service/internal/domain/services"
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

	usecase := usecases.NewGetPayoutByIDUsecase(h.payoutService)

	payout, err := usecase.Execute(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Payout not found"},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payout": payout})
}

func (h *PayoutHandler) CreatePayout(ctx *gin.Context) {
	usecase := usecases.NewCreatePayoutUsecase(h.payoutService)

	var requestBody dtos.CreatePayoutRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	payout, err := usecase.Execute(ctx, &requestBody)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Payout created successfully",
			"payout":  &payout,
		},
	)
}

func (h *PayoutHandler) UpdatePayoutDescription(ctx *gin.Context) {
	usecase := usecases.NewUpdatePayoutDescriptionUsecase(h.payoutService)

	var requestBody dtos.UpdatePayoutDescriptionRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	payout, err := usecase.Execute(ctx, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "Payout description updated successfully",
			"payout":  &payout,
		},
	)
}
