package services

import (
	"context"

	"github.com/keyinvoker/go-payout-service/internal/domain/models"
	"github.com/keyinvoker/go-payout-service/internal/domain/repositories"
)

type PayoutService struct {
	payoutRepo *repositories.PayoutRepository
}

func NewPayoutService(payoutRepo *repositories.PayoutRepository) *PayoutService {
	return &PayoutService{payoutRepo: payoutRepo}
}

func (service *PayoutService) GetByID(ctx context.Context, id int) (*models.Payout, error) {
	return service.payoutRepo.GetByID(ctx, id)
}
