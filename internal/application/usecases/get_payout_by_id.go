package usecases

import (
	"context"

	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/models"
)

type GetPayoutByIDUsecase struct {
	payoutService *services.PayoutService
}

func NewGetPayoutByIDUsecase(payoutService *services.PayoutService) *GetPayoutByIDUsecase {
	return &GetPayoutByIDUsecase{payoutService: payoutService}
}

func (u *GetPayoutByIDUsecase) Execute(ctx context.Context, id int) (*models.Payout, error) {
	payout, err := u.payoutService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return payout, nil
}
