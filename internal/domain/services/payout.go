package services

import (
	"context"

	"github.com/keyinvoker/go-payout-service/internal/domain/models"
	repositories "github.com/keyinvoker/go-payout-service/internal/domain/repositories/database"
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

func (service *PayoutService) Create(
	ctx context.Context,
	payout *models.Payout,
) (*models.Payout, error) {

	err := service.payoutRepo.Create(ctx, payout)
	if err != nil {
		return nil, err
	}

	return payout, nil
}

func (service *PayoutService) UpdateDescription(
	ctx context.Context,
	id int,
	description string,
) (*models.Payout, error) {

	payout, err := service.payoutRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	payout.Description = &description

	err = service.payoutRepo.Update(ctx, payout)
	if err != nil {
		return nil, err
	}

	return payout, nil
}
